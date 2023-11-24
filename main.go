package main

import (
	"github.com/gin-gonic/gin"

	"simple-api/src/routes"
	"simple-api/src/models"
)

func main() {
	InitEnv()

	models.MigrateUser()

	r := gin.Default()

	routes.Students(r)
	routes.Users(r)
	routes.Authentications(r)

	r.Run()
}