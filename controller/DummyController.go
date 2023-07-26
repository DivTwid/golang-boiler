package controller

import (
	"net/http"

	"github.com/DivTwid/golang-boiler/dto"
	"github.com/DivTwid/golang-boiler/service"

	"github.com/gin-gonic/gin"
)

type DummyController struct {
	ds service.DummyService
}

func NewDummyController() *DummyController {
	return &DummyController{
		ds: service.NewDummyService(),
	}
}

func (d DummyController) GetVal(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, d.ds.GetVal())
}

func (d DummyController) AddVal(ctx *gin.Context) {
	bind := dto.UserDto{}
	ctx.ShouldBindJSON(&bind)
	ctx.JSON(http.StatusCreated, d.ds.AddVal(bind))
}
