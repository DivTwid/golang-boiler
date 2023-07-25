package seeders

import (
	"log"

	"github.com/DivTwid/golang-boiler/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func genRandUser() model.User {
	user := model.User{
		Name:    uuid.NewString()[:10],
		Email:   uuid.NewString()[:5] + "@" + uuid.NewString()[:5] + ".com",
		PhoneNo: uuid.NewString()[:10],
	}
	return user
}

func UserSeeder(db *gorm.DB) []model.User {
	user := []model.User{}
	for i := 0; i < 5; i++ {
		usr, err := model.CreateUser(db, genRandUser())
		if err != nil {
			log.Print("Error", err.Error())
		} else {
			user = append(user, usr)
		}
	}
	return user
}
