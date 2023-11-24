package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"simple-api/src/database"
	"simple-api/src/models"
	"simple-api/src/utils"
)

func PostHandler(c *gin.Context) {
	var student models.Student
	var headers models.Headers

	c.BindHeader(&headers)

	tokenStr, _ := strings.CutPrefix(headers.Authorization, "Bearer ")

	token := utils.JWT{ TokenStr: tokenStr, Key: []byte(os.Getenv("ACCESS_TOKEN_SECRET"))}

	claims, err := token.VerifyToken()

	if err != nil {
		utils.ErrorMessage(c, http.StatusInternalServerError, "internal server error")
		return
	}

	fmt.Println(claims)

	if c.Bind(&student) == nil {
		result := database.DB().Create(&student)

		if result.Error != nil {
			utils.ErrorMessage(c, http.StatusInternalServerError, "internal server error")
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "student added"})
	}
}

func GetAllHandler(c *gin.Context) {
	var students []models.Student

	result := database.DB().Find(&students)

	if result.Error != nil {
		utils.ErrorMessage(c, http.StatusInternalServerError, "internal server error")
		return
	}

	if students == nil {
		c.JSON(http.StatusOK, gin.H{"message": "data empty"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": students})
}

func GetByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var student models.Student

	result := database.DB().First(&student, id)

	if result.Error != nil {
		if student.Student_id == 0 {
			utils.ErrorMessage(ctx, http.StatusNotFound, "id not found")
			return
		}

		utils.ErrorMessage(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": student})
}

func UpdateByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var student models.Student

	result := database.DB().First(&student, id)

	if result.Error != nil {
		if student.Student_id == 0 {
			utils.ErrorMessage(ctx, http.StatusNotFound, "id not found")
			return
		}

		utils.ErrorMessage(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	if ctx.Bind(&student) == nil {
		if database.DB().Save(&student).Error != nil {
			utils.ErrorMessage(ctx, http.StatusInternalServerError, "internal server error")
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "student updated"})
	}
}

func DeleteByIdHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var student models.Student

	result := database.DB().First(&student, id)

	if result.Error != nil {
		if student.Student_id == 0 {
			utils.ErrorMessage(ctx, http.StatusNotFound, "id not found")
			return
		}

		utils.ErrorMessage(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	deleteResult := database.DB().Delete(&student)

	if deleteResult.Error != nil {
		utils.ErrorMessage(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "student deleted"})
}