package seeders

import (
	"simple-api/src/database"
	"simple-api/src/models"
)

func SeedStudents(data []models.Student) error {
	return database.DB().Create(&data).Error
}