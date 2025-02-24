package main

import (
	"StudentManagementAdvance/config"
	"StudentManagementAdvance/routes"
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	dbConfig, err := config.Init("config.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("Loaded DB Config: %+v\n", dbConfig)
	fmt.Printf("MongoDB URI: %s\n", dbConfig.MongoConnection.ServerIP)
	fmt.Printf("MongoDB Database: %s\n", dbConfig.MongoConnection.Database)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Change this based on your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.UseRoutes(router)

	fmt.Println("Server running on port 3000")

	router.Run(":3000")
}

// TODO: Cyclic error Golang
