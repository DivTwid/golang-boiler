package service

import (
	"github.com/DivTwid/golang-boiler/dto"
	"github.com/DivTwid/golang-boiler/model"
)

type DummyService interface {
	GetVal() string
	AddVal(user dto.UserDto) model.User
}
type dummyService struct {
}

func NewDummyService() DummyService {
	return &dummyService{}
}

func (ds dummyService) GetVal() string {
	return "pong"
}

func (ds dummyService) AddVal(user dto.UserDto) model.User {
	userModel := model.User{
		Name:    user.Name,
		Email:   user.Email,
		PhoneNo: user.PhoneNo,
	}

	// ret, err := model.CreateUser(config.PqDB, userModel)
	// if err != nil {
	// 	log.Fatal("Error while creating user", err)
	// 	return userModel
	// }
	return userModel
}
