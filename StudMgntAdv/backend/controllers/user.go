package controllers

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"StudentManagementAdvance/config"
	"StudentManagementAdvance/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection
var examCollection *mongo.Collection
var examApplicationsCollection *mongo.Collection

func InitController() {
	if config.DB == nil {
		panic("Database not connected")
	}
	userCollection = config.DB.Collection("users")
	examCollection = config.DB.Collection("examForm")
	examApplicationsCollection = config.DB.Collection("examApplications")
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func RegisterUser(ctx *gin.Context) {
	var user model.Student

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return

	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	user.Password = hashedPassword
	user.CreatedOn = time.Now()
	user.UpdatedOn = time.Now()
	user.IsActive = true
	user.CreatedBy = "Admin"

	var existingUser model.Student
	err = userCollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "User already present"})
		return
	}

	_, err = userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving user to database"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

var secretKey = []byte("secretKey")

func generateToken(email, roles string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"roles": roles,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, error := token.SignedString(secretKey)
	if error != nil {
		return "", error
	}
	model.SaveToken(tokenString, email)
	return tokenString, nil
}
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.GetHeader("Authorization")

		if tokenHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			ctx.Abort()
			return
		}

		const bearerPrefix = "Bearer "
		if len(tokenHeader) <= len(bearerPrefix) || tokenHeader[:len(bearerPrefix)] != bearerPrefix {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			ctx.Abort()
			return
		}

		tokenString := tokenHeader[len(bearerPrefix):]

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		email, ok := claims["email"].(string)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token data"})
			ctx.Abort()
			return
		}

		ctx.Set("email", email)
		ctx.Next()
	}
}

func LoginUser(ctx *gin.Context) {
	var user model.Student
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	var existingUser model.Student
	err := userCollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "User not found"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}

	token, err := generateToken(existingUser.Email, existingUser.Roles)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token, "roles": existingUser.Roles})
}

func UserProfile(ctx *gin.Context) {
	email, exists := ctx.Get("email")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}
	var user model.Student
	error := userCollection.FindOne(context.TODO(), bson.M{"email": email.(string)}).Decode(&user)
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user profile"})
		return
	}
	ctx.JSON(http.StatusOK, user)

}

func AdminProfile(ctx *gin.Context) {
	email, exists := ctx.Get("email")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}
	var adminProfile model.Student
	error := userCollection.FindOne(context.TODO(), bson.M{"email": email.(string)}).Decode(&adminProfile)
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user profile"})
		return
	}
	ctx.JSON(http.StatusOK, adminProfile)
}

func ViewAllUsers(ctx *gin.Context) {
	_, exists := ctx.Get("email")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}
	var users []model.Student
	cursor, error := userCollection.Find(context.TODO(), bson.M{})
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch users"})
		return
	}
	defer cursor.Close(context.TODO())
	if error = cursor.All(context.TODO(), &users); error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode users"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func DeleteUser(ctx *gin.Context) {

	email := ctx.Param("email")

	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	result, err := userCollection.DeleteOne(context.TODO(), bson.M{"email": email})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	if result.DeletedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func GetUserForUpdate(ctx *gin.Context) {
	email := ctx.Param("email")

	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	var user model.Student
	err := userCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func UpdateUser(ctx *gin.Context) {

	_, exists := ctx.Get("email")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Missing token"})
		return
	}

	email := ctx.Param("email")
	if email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}
	var existingUser struct {
		CreatedOn time.Time `bson:"created_on"`
	}
	err := userCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&existingUser)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var updatedData struct {
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Roles     string `json:"roles"`
		IsActive  bool   `json:"is_active"`
		Password  string `json:"password,omitempty"` 
	}

	if err := ctx.ShouldBindJSON(&updatedData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	updateFields := bson.M{
		"username":   updatedData.Username,
		"first_name": updatedData.FirstName,
		"last_name":  updatedData.LastName,
		"roles":      updatedData.Roles,
		"is_active":  updatedData.IsActive,
		"created_on": existingUser.CreatedOn, 
		"updated_on": time.Now(),             
	}
	if updatedData.Password != "" {
		hashedPassword, err := hashPassword(updatedData.Password)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
			return
		}
		updateFields["password"] = hashedPassword
	}

	filter := bson.M{"email": email}
	update := bson.M{"$set": updateFields}

	result, err := userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	if result.MatchedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func GenerateRandomID() string {
	bytes := make([]byte, 8)
	_, err := rand.Read(bytes)
	if err != nil {
		return "" 
	}
	return hex.EncodeToString(bytes)
}

func CreateExam(ctx *gin.Context) {
	emailFromToken, exists := ctx.Get("email")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Missing token"})
		return
	}

	var exam model.Exam
	if err := ctx.ShouldBindJSON(&exam); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	exam.ExamID = GenerateRandomID()
	exam.CreatedBy = emailFromToken.(string)
	exam.CreatedAt = time.Now()

	_, err := examCollection.InsertOne(context.TODO(), exam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create exam"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Exam created successfully", "examId": exam.ExamID})
}
func GetExam(ctx *gin.Context) {
	_, exists := ctx.Get("email")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Missing token"})
		return
	}
	cursor, err := examCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch exams"})
		return
	}
	defer cursor.Close(context.TODO())

	var exams []model.Exam
	if err := cursor.All(context.TODO(), &exams); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode exams"})
		return
	}

	ctx.JSON(http.StatusOK, exams)

}

func DeleteExam(ctx *gin.Context) {
	_, exists := ctx.Get("email")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Missing token"})
		return
	}

	id := ctx.Param("id")                
	fmt.Println("Received Exam ID:", id) 

	result, err := examCollection.DeleteOne(context.TODO(), bson.M{"exam_id": id})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete exam"})
		return
	}

	if result.DeletedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Exam not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Exam deleted successfully"})
}

func UpdateExam(ctx *gin.Context) {
	_, exists := ctx.Get("email")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Missing token"})
		return
	}

	id := ctx.Param("id")                           
	fmt.Println("Received Exam ID for update:", id) 

	var updatedExam model.Exam
	if err := ctx.ShouldBindJSON(&updatedExam); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	filter := bson.M{"exam_id": id}
	update := bson.M{
		"$set": bson.M{
			"exam_name":              updatedExam.ExamName,
			"exam_date":              updatedExam.ExamDate,
			"exam_duration":          updatedExam.Duration,
			"exam_type":              updatedExam.ExamType,
			"application_close_date": updatedExam.ApplicationCloseDate,
		},
	}

	result, err := examCollection.UpdateOne(context.TODO(), filter, update, options.Update().SetUpsert(false))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update exam"})
		return
	}

	if result.MatchedCount == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Exam not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Exam updated successfully"})
}

func ApplyForExam(ctx *gin.Context) {
    email, exists := ctx.Get("email")
    if !exists {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Missing token"})
        return
    }

    examID := ctx.Param("id") 
    if examID == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Exam ID is required"})
        return
    }

    var exam model.Exam
    filter := bson.M{"exam_id": examID}
    err := examCollection.FindOne(context.TODO(), filter).Decode(&exam)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "Exam not found"})
        return
    }

    if exam.ApplicationCloseDate.Before(time.Now()) {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Application deadline has passed"})
        return
    }

    applicationFilter := bson.M{"exam_id": examID, "email": email}
    var existingApplication model.ExamApplication
    err = examApplicationsCollection.FindOne(context.TODO(), applicationFilter).Decode(&existingApplication)
    if err == nil {
        ctx.JSON(http.StatusConflict, gin.H{"error": "You have already applied for this exam"})
        return
    }

    newApplication := model.ExamApplication{
        ExamID:    examID,
        ExamName:  exam.ExamName,
        Email:     email.(string),
        AppliedAt: time.Now(),
		
    }

    _, err = examApplicationsCollection.InsertOne(context.TODO(), newApplication)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to apply for the exam"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Successfully applied for the exam"})
}
func GetApplicationsHandler(ctx *gin.Context) {
	_, exists := ctx.Get("email")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Missing token"})
		return
	}

	examId := ctx.Param("examId")

	fmt.Println("Fetching applications for Exam ID:", examId)

	filter := bson.M{"exam_id": examId}

	cursor, err := examApplicationsCollection.Find(context.TODO(), filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch applications"})
		return
	}
	defer cursor.Close(context.TODO())

	var applications []model.ExamApplication
	if err := cursor.All(context.TODO(), &applications); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding applications"})
		return
	}

	fmt.Println("Applications found:", applications)

	if applications == nil {
		applications = []model.ExamApplication{}
	}

	ctx.JSON(http.StatusOK, applications)
}
