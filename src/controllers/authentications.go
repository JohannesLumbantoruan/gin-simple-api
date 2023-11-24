package controllers

import (
	"os"
	"net/http"
	"simple-api/src/database"
	"simple-api/src/models"
	"simple-api/src/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func PostAuthenticationHandler(c *gin.Context) {
	var credential models.Credential
	var user models.User

	if c.Bind(&credential) == nil {
		result := database.DB().Where(&models.User{ Username: credential.Username }).First(&user)

		if result.Error != nil {
			if user.Id == 0 {
				utils.ErrorMessage(c, http.StatusNotFound, "user not found")
				return
			}

			utils.ErrorMessage(c, http.StatusInternalServerError, "internal server error")
			return
		}

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credential.Password))

		if err != nil {
			utils.ErrorMessage(c, http.StatusUnauthorized, "credential wrong")
			return
		}

		claims := jwt.MapClaims{
			"exp": time.Now().Add(time.Minute * 1).Unix(),
			"userId": user.Id,
			"username": user.Username,
			"fullname": user.Fullname,
		}

		// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		accessToken := utils.JWT{ Claims: claims, Key: []byte(os.Getenv("ACCESS_TOKEN_SECRET"))}
		refreshToken := utils.JWT{ Claims: claims, Key: []byte(os.Getenv("REFRESH_TOKEN_SECRET"))}

		// accessToken, _ := token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
		// refreshToken, _ := token.SignedString([]byte(os.Getenv("REFRESH_TOKEN_SECRET")))

		c.JSON(http.StatusOK, gin.H{ "data": gin.H {
			"accessToken": accessToken.GenerateToken(),
			"refreshToken": refreshToken.GenerateToken(),
		}})

		return
	}

	utils.ErrorMessage(c, http.StatusInternalServerError, "internal server error")
}