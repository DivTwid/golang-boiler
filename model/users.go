package model

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Name    string
	Email   string `json:"email"`
	PhoneNo string
}

func CreateUser(db *gorm.DB, user User) (User, error) {
	result := db.Create(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func GetUsers(db *gorm.DB) []User {
	var user []User
	result := db.Find(&user)
	if result.Error != nil {
		fmt.Println("result error", result.Error)
	}
	return user
}
