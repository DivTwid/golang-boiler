package controller

import (
	"net/http"

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
	ctx.String(http.StatusOK, d.ds.GetVal())
	return
}

func (d DummyController) AddVal(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, d.ds.AddVal())
}
