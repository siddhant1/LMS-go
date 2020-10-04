package routes

import (
	controller "ilearn-backend/controllers"

	"github.com/gin-gonic/gin"
)

// TeacherRoute exported
func TeacherRoute(router *gin.Engine) {
	teacher := router.Group("api/v1/teacher")
	{
		teacher.POST("/", controller.AddTeacher)
		teacher.GET("/", controller.GetAllTeachers)
	}
}
