package main

import (
	"sql-connection/config"
	"sql-connection/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	router := gin.Default()
	router.Use(cors.Default())
	routes.UseRoutes(router)
	router.Run(":3000")
}
