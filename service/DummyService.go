package service

import (
	"fmt"
	"log"

	"github.com/DivTwid/golang-boiler/config"
	"github.com/DivTwid/golang-boiler/dto"
	"github.com/DivTwid/golang-boiler/model"
	"gorm.io/gorm"
)

type DummyService interface {
	GetVal() string
	AddVal(user dto.UserDto) model.User
}
type dummyService struct {
	db *gorm.DB
}

func NewDummyService() DummyService {
	return &dummyService{
		db: config.PqDB,
	}
}

func (ds dummyService) GetVal() string {
	return "pong"
}

func (ds dummyService) AddVal(user dto.UserDto) model.User {
	fmt.Println("user", user)
	userModel := model.User{
		Name:    user.Name,
		Email:   user.Email,
		PhoneNo: user.PhoneNo,
	}
	ret, err := model.CreateUser(ds.db, userModel)
	if err != nil {
		log.Fatal("Error while creating user", err)
	}
	return ret
}
