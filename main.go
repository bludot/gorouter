package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bludot/gorouter/controller"
	router "github.com/bludot/gorouter/router"
	"github.com/bludot/gorouter/router/entities"
)

type RootController struct {
	controller.Controller
}

func (c *RootController) Run(ctx context.Context, params *entities.RouteParams, queryParams *entities.QueryParams) error {
	log.Println("Controller:", c.Name)
	log.Println("Params:", params)
	return nil
}

func NewRootController() controller.IController {
	thisController := &controller.Controller{
		Name: "RootController",
	}
	return &RootController{
		Controller: *thisController,
	}
}

func main() {
	mainRouter := router.NewRouter()
	mainRouter.AddRoute(router.Route{
		Controller: NewRootController(),
		Path:       "/",
	})
	mainRouter.AddRoute(router.Route{
		Controller: NewRootController(),
		Path:       "/test/$id",
	})
	http.Handle("/", mainRouter)
	http.ListenAndServe(":8080", nil)
}
