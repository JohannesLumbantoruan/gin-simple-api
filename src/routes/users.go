package routes

import (
	"simple-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func Users(r *gin.Engine) *gin.RouterGroup {
	users := r.Group("/users")

	users.POST("/", controllers.PostUserHandler)

	return users
}