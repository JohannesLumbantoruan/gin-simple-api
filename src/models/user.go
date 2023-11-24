package models

import (
	"simple-api/src/database"
)

type User struct {
	Id	uint `json:"id" binding:"required" gorm:"primaryKey"`
	Username	string `json:"username" binding:"required"`
	Password	string `json:"password" binding:"required"`
	Fullname	string `json:"fullname" binding:"required"`
}

func MigrateUser() {
	var user User
	database.DB().AutoMigrate(&user)
}