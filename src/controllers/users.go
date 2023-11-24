package controllers

import (
	"simple-api/src/database"
	"simple-api/src/models"
	"simple-api/src/utils"

	"github.com/gin-gonic/gin"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

func PostUserHandler(c *gin.Context) {
	var user models.User

	if c.Bind(&user) == nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

		if err != nil {
			utils.ErrorMessage(c, http.StatusInternalServerError, "internal server error")
			return
		}

		user.Password = string(hashedPassword)

		if result := database.DB().Create(&user).Error; result != nil {
			utils.ErrorMessage(c, http.StatusInternalServerError, "internal server error")
			return
		}

		c.JSON(http.StatusOK, gin.H{ "message": "user created" })
	}
}