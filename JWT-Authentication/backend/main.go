package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ChangePassword struct {
    OldPassword     string `json:"oldPassword" binding:"required"`
    NewPassword     string `json:"newPassword" binding:"required"`
    ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

const uri = "mongodb+srv://bhosalerupesh67:L2hSiWX3IeytWQNm@jwt-users.8jhul.mongodb.net/?retryWrites=true&w=majority&appName=JWT-Users"

var mongoClient *mongo.Client
var userCollection *mongo.Collection

func init() {
	if err := connect_to_mongo(); err != nil {
		fmt.Println("could not connect to mongo")
	} else {
		fmt.Println("connected to mongo")
	}
	userCollection = mongoClient.Database("users").Collection("users")
}

func connect_to_mongo() error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return err
	}
	err = client.Ping(context.TODO(), nil)
	mongoClient = client
	return err
}

var secretKey = []byte("secretKey")

func generateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(secretKey)
}

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")

        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
            c.Abort()
            return
        }

        token, err := jwt.Parse(tokenString[7:], func(token *jwt.Token) (interface{}, error) {
            return secretKey, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
            c.Abort()
            return
        }

        c.Next()
    }
}

func main() {
	router := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}

	router.Use(cors.New(config))

	router.POST("/addUser", func(ctx *gin.Context) {
		var user User
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//  check if user already exists
		var existingUser User
		err := userCollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)
		if err == nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
			return
		} else if err != mongo.ErrNoDocuments {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		// hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		user.Password = string(hashedPassword)

		// add user
		insertUser, error := userCollection.InsertOne(context.TODO(), user)
		if error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": error})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message":    "User created successfully",
			"insertedId": insertUser.InsertedID,
		})
	})

	router.POST("/login", func(ctx *gin.Context) {
		var user User
		if error := ctx.ShouldBind(&user); error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": error})
			return
		}
		var existingUser User
		error := userCollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)
		if error == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "user not found"})
			return
		} else if error != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": error})
			return
		}

		error = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
		if error != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "password not match"})
			return
		}
		token, err := generateToken(existingUser.Email)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Login successful","token":token})

	})

	router.GET("/profile", AuthMiddleware(), func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		token, _ := jwt.Parse(tokenString[7:], func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
	
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email := claims["email"].(string)
	
			var user User
			err := userCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
			if err != nil {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
	
			ctx.JSON(http.StatusOK, gin.H{
				"name": user.Name, 
				"email": user.Email,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		}
	})
	router.POST("/changepassword", AuthMiddleware(), func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		fmt.Println("Received token:", tokenString)
	
		token, _ := jwt.Parse(tokenString[7:], func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
	
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email := claims["email"].(string)
			fmt.Println("Email extracted from token:", email)
	
			var changePassword ChangePassword
			err := ctx.BindJSON(&changePassword)
			if err != nil {
				fmt.Println("Error parsing JSON:", err)
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
				return
			}
			fmt.Println("Received change password request:", changePassword)
	
			var existingUser User
			err = userCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&existingUser)
			if err != nil {
				fmt.Println("User not found:", err)
				ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
				return
			}
	
			err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(changePassword.OldPassword))
			if err != nil {
				fmt.Println("Old password does not match:", err)
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Old password not match"})
				return
			}
	
			if changePassword.NewPassword != changePassword.ConfirmPassword {
				fmt.Println("New passwords do not match")
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "New password and confirm password not match"})
				return
			}
	
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(changePassword.NewPassword), bcrypt.DefaultCost)
			if err != nil {
				fmt.Println("Error hashing password:", err)
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}
	
			updateResult, err := userCollection.UpdateOne(context.TODO(), bson.M{"email": email}, bson.M{"$set": bson.M{"password": string(hashedPassword)}})
			if err != nil {
				fmt.Println("Database update error:", err)
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}
			fmt.Println("Update result:", updateResult)
	
			ctx.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
		} else {
			fmt.Println("Invalid token")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		}
	})

	router.GET("/users",func(ctx *gin.Context) {
		var users []User
		cursor, err := userCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		defer cursor.Close(context.TODO())
		for cursor.Next(context.TODO()) {
			var user User
			cursor.Decode(&user)
			users = append(users, user)
		}
		ctx.JSON(http.StatusOK, gin.H{"users": users})
	})

	router.DELETE("/deleteUser/:email",func(ctx *gin.Context) {
		email := ctx.Param("email")
		updateResult, err := userCollection.DeleteOne(context.TODO(), bson.M{"email": email})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully","deletedCount":updateResult.DeletedCount})
	})
	

	router.Run(":3000")
}
