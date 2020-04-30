package main

import (
	"log"

	api "github.com/k-ueki/AGPlus/server/adaptor/api/controller"

	"github.com/gin-gonic/gin"
)

type (
	App struct {
		ClassGetController api.ClassGetController
	}
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	router := gin.Default()
	setRoutes(router)
	return router.Run(":8888")
}

func setRoutes(r *gin.Engine) {
	app := &App{}
	route := r.Group("/")
	//auth := r.AuthorizeWrapper

	{
		classes := route.Group("classes")
		classes.GET("/", app.ClassGetController.Hoge)
	}
}
