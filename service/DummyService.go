package service

import (
	"log"

	"github.com/DivTwid/golang-boiler/db"
	"github.com/DivTwid/golang-boiler/dto"
	"github.com/DivTwid/golang-boiler/model"
)

type DummyService interface {
	GetVal() []model.User
	AddVal(user dto.UserDto) model.User
}
type dummyService struct {
}

func NewDummyService() DummyService {
	return &dummyService{}
}

func (ds dummyService) GetVal() []model.User {
	ret := model.GetUsers(db.PqDB)
	return ret
}

func (ds dummyService) AddVal(user dto.UserDto) model.User {
	userModel := model.User{
		Name:    user.Name,
		Email:   user.Email,
		PhoneNo: user.PhoneNo,
	}

	ret, err := model.CreateUser(db.PqDB, userModel)
	if err != nil {
		log.Fatal("Error while creating user", err)
		return userModel
	}
	return ret
}
