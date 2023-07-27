package model

import (
	"fmt"

	"github.com/DivTwid/golang-boiler/logs"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Name    string
	Email   string `gorm:"unique" json:"email"`
	PhoneNo string
}

func CreateUser(db *gorm.DB, user User) (User, error) {
	result := db.Create(&user)
	if result.Error != nil {
		logs.Error("error while generating user" + result.Error.Error())
		return User{}, result.Error
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
