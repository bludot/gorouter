package test

import (
	"context"
	controller "github.com/bludot/gorouter/core/controller"
	"github.com/bludot/gorouter/core/renderer"
	"github.com/bludot/gorouter/core/router/entities"
	"log"
)

type TestController struct {
	controller.Controller
}

func (c *TestController) Handle(ctx context.Context, params *entities.RouteParams, queryParams *entities.QueryParams) error {
	log.Println("Controller:", c.Name)
	log.Println("Params:", params)
	renderer.GetRender().Render("index.html", map[string]string{
		"body": params.Get("id"),
	})
	return nil
}

func NewTestController() controller.IController {
	thisController := &controller.Controller{
		Name: "RootController",
	}
	return &TestController{
		Controller: *thisController,
	}
}
