package main

// type Student struct {
// 	FirstName string `json:"firstName"`
// 	LastName  string `json:"lastName"`
// 	EmailId   string `json:"emailId"`
// }

// var StudentData []Student
/*
Vuejs
- Student page:
	a. Insert new student
	b. Get all student list
	c. Delete student
	d. Update student info
Golang
- Student handlers:
	a. Create the endpoints for the above operations
	b. Get Student data functionality implementation
	c. Delete / Remove student info
	d. Update

	Note:
	Validations
	Method Type
	-GET
	-POST
	-PUT
	-Delete

Profile pic update (optional)
- Photo upload


*/

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Payload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	EmailId   string `json:"emailId"`
}
	
func postData(c *gin.Context) {
	var pd Payload
	if err := c.BindJSON(&pd); err != nil {
		fmt.Println("Error while binding payload data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if pd.FirstName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "First name is required"})
		return
	}

	fmt.Println("First name:", pd.FirstName)
	fmt.Println("Last name:", pd.LastName)
	fmt.Println("Email id:", pd.EmailId)

	c.JSON(200, gin.H{
		"message":     "Data received successfully",
		"payloadData": pd, // Print the received data
	})
}

func main() {
	route := gin.Default()
	route.Use(cors.Default())

	route.POST("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": []gin.H{
				{
					"id":   "1234",
					"name": "Rupesh Bhosale",
				},
				{
					"id":   "5678",
					"name": "Reshma Bhosale",
				},
			},
		})
	})

	route.POST("/postdata", postData)

	route.Run(":3000")
}
