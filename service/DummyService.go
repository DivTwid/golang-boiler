package service

import (
	"fmt"

	"github.com/DivTwid/golang-boiler/db"
	"github.com/DivTwid/golang-boiler/dto"
	"github.com/DivTwid/golang-boiler/model"
)

type DummyService interface {
	GetVal() []model.User
	AddVal(user dto.UserDto) (model.User, error)
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

func (ds dummyService) AddVal(user dto.UserDto) (model.User, error) {
	userModel := model.User{
		Name:    user.Name,
		Email:   user.Email,
		PhoneNo: user.PhoneNo,
	}

	ret, err := model.CreateUser(db.PqDB, userModel)
	if err != nil {
		fmt.Println("error while user create")
		return userModel, err
	}
	return ret, nil
}
