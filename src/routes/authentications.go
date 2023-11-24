package routes

import (
	"simple-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func Authentications(r *gin.Engine) *gin.RouterGroup {
	authentications := r.Group("/authentications")

	authentications.POST("/", controllers.PostAuthenticationHandler)

	return authentications
}