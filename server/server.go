package main

import (
	"log"

	"github.com/gin-contrib/cors"

	"github.com/jinzhu/gorm"
	"github.com/k-ueki/AGPlus/server/common"

	"github.com/gin-gonic/gin"

	api "github.com/k-ueki/AGPlus/server/adaptor/api/controller"
)

type (
	App struct {
		ClassGetController api.ClassGetController
	}

	Server struct {
		DB     *gorm.DB
		Router *gin.Engine
	}
)

func NewServer() (*Server, error) {
	db, err := common.NewDB()
	if err != nil {
		return nil, err
	}
	return &Server{
		DB:     db,
		Router: gin.Default(),
	}, nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	server, err := NewServer()
	if err != nil {
		return err
	}

	server.Router.Use(cors.New(cors.Config{AllowOrigins: []string{"*"}}))

	setRoutes(server)

	return server.Router.Run(":8888")
}

func setRoutes(r *Server) {
	app := &App{
		ClassGetController: *api.NewClassGetController(r.DB),
	}
	router := r.Router.Group("/")
	//auth := r.AuthorizeWrapper

	{
		classes := router.Group("classes")
		classes.GET("/", app.ClassGetController.List)
	}
}
