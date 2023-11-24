package routes

import (
	"github.com/gin-gonic/gin"

	"simple-api/src/controllers"
)

func Students(r *gin.Engine) *gin.RouterGroup {
	students := r.Group("/students")

	students.POST("/", controllers.PostHandler)

	students.GET("/", controllers.GetAllHandler)

	students.GET("/:id", controllers.GetByIdHandler)

	students.PUT("/:id", controllers.UpdateByIdHandler)

	students.DELETE("/:id", controllers.DeleteByIdHandler)

	return students
}