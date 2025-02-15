package routes

import (
	"StudentManagementAdvance/controllers"

	"github.com/gin-gonic/gin"
)

func UseRoutes(r *gin.Engine) {
	controllers.InitController()
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)
	// creates routes
	restricted := r.Group("/")
	restricted.Use(controllers.AuthMiddleware())
	{
		restricted.GET("/userProfile", controllers.UserProfile)
		restricted.GET("/adminProfile",controllers.AdminProfile)
		restricted.GET("/viewAllUsers",controllers.ViewAllUsers)
		restricted.DELETE("/deleteUser/:email",controllers.DeleteUser)
		restricted.GET("/getUserForUpdate/:email", controllers.GetUserForUpdate)
		restricted.PUT("/updateUser/:email",controllers.UpdateUser)
		restricted.POST("/createExam",controllers.CreateExam)
		restricted.GET("/getExams",controllers.GetExam)
		restricted.DELETE("/deleteExam/:id",controllers.DeleteExam)
		restricted.PUT("/updateExam/:id",controllers.UpdateExam)
		restricted.POST("/applyExams/:id", controllers.ApplyForExam)
		restricted.GET("/getApplications/:examId", controllers.GetApplicationsHandler)

	}
}
