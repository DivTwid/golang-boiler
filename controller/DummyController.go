package controller

import (
	"net/http"

	"github.com/DivTwid/golang-boiler/dto"
	"github.com/DivTwid/golang-boiler/service"
	"github.com/DivTwid/golang-boiler/utils"

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
	result, err := d.ds.AddVal(bind)
	if err != nil {
		utils.HandleError(ctx, 500, err.Error())
	} else {
		ctx.JSON(http.StatusCreated, result)
	}
}
