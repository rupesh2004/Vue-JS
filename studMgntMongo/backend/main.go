package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/exp/rand"
)

type Student struct {
	StudentID int    `json:"studentID"`
	FullName  string `json:"fullName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required"`
	Dob       string `json:"dob"`
}

func generateStudentID() int {
	return rand.Intn(9000) + 1000
}

var mongoClient *mongo.Client
var studentCollection *mongo.Collection

const uri = "mongodb+srv://bhosalerupesh67:yjk7k48j3k4R2Zde@studentsdata.151co.mongodb.net/?retryWrites=true&w=majority&appName=studentsData"

func init() {
	if err := connect_to_mongo(); err != nil {
		log.Fatal("Could not connect to MongoDB")
	} else {
		fmt.Println("Connected to MongoDB")
	}
	studentCollection = mongoClient.Database("studentsData").Collection("students")
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

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", func(ctx *gin.Context) {
		// fetch all data
		cursor, err := studentCollection.Find(context.TODO(), bson.D{})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not fetch students"})
			return
		}
		defer cursor.Close(context.TODO())
		var students []Student
		for cursor.Next(context.TODO()) {
			var student Student
			if err = cursor.Decode(&student); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not decode student"})
				return
			}
			students = append(students, student)
		}
		if err = cursor.Err(); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not fetch students"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"students": students})
	})

	router.POST("/studData", func(ctx *gin.Context) {
		var student Student
		if err := ctx.ShouldBindJSON(&student); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		student.StudentID = generateStudentID()
		_, err := studentCollection.InsertOne(context.TODO(), student)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Could not insert student"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Student inserted", "studentID": student.StudentID})
	})

	router.DELETE("/deleteStudent/:studentID", func(ctx *gin.Context) {
		studentID := ctx.Param("studentID")

		// Convert studentID to int
		studentIDInt, err := strconv.Atoi(studentID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid studentID"})
			return
		}

		studentIDBson := bson.D{{"studentid", studentIDInt}}

		result, err := studentCollection.DeleteOne(context.TODO(), studentIDBson)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete student", "details": err.Error()})
			return
		}

		if result.DeletedCount == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Student %d deleted successfully", studentIDInt)})
	})

	router.PUT("/updateStudent/:studentID", func(ctx *gin.Context) {
		studentID := ctx.Param("studentID")
		studentIDInt, err := strconv.Atoi(studentID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid studentID"})
			return
		}
		studentIDBson := bson.D{{"studentid", studentIDInt}}

		var student Student
		err = studentCollection.FindOne(context.TODO(), studentIDBson).Decode(&student)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
			return
		}
		if err := ctx.ShouldBindJSON(&student); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err = studentCollection.UpdateOne(context.TODO(), studentIDBson, bson.D{{"$set", student}})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update student", "details": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Student %d updated successfully", studentIDInt)})


	})

	router.Run(":3000")
}
