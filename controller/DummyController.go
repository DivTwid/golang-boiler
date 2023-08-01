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

// GetVals 		 godoc
// @Summary      List accounts
// @Description  get accounts
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /api/GetVal [get]
func (d DummyController) GetVal(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, d.ds.GetVal())
}

// AddVals 		 godoc
// @Summary      Add Values
// @Description  get accounts
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /api/addVals [post]
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
