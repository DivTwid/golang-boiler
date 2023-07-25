package model

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Name    string
	Email   string `gorm:"uniqueKey;uniqueIndex" json:"email"`
	PhoneNo string
}

func CreateUser(db *gorm.DB, user User) (User, error) {
	result := db.Create(&user) // pass pointer of data to Create
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func GetUsers(db *gorm.DB) {
	var user User
	result := db.Find(&user)
	fmt.Println("result", result)
}
