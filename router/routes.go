package router

import (
	"net/http"

	"github.com/DivTwid/golang-boiler/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	routes := gin.Default()
	routes.LoadHTMLGlob("templates/*")
	routes.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Go Boiler",
			"para":  "Some details will go here",
		})
	})

	//Controller Call
	dummyController := controller.NewDummyController()
	DummyRoutes(routes, *dummyController)
	return routes
}

func DummyRoutes(route *gin.Engine, ctrl controller.DummyController) {
	route.GET("/addVal", ctrl.AddVal)
	route.GET("/getVal", ctrl.GetVal)
}
