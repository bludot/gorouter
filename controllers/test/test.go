package test

import (
	"context"
	controller "github.com/bludot/gorouter/core/controller"
	"github.com/bludot/gorouter/core/renderer"
	"github.com/bludot/gorouter/core/router/entities"
	"log"
	"net/http"
)

type TestController struct {
	controller.Controller
}

func (c *TestController) TestRoute(ctx context.Context, params *entities.RouteParams, queryParams *entities.QueryParams) error {
	log.Println("Handler:", c.Name)
	log.Println("Params:", params)
	renderer.GetRender().Render("index.html", map[string]string{
		"body": params.Get("id"),
	}, http.StatusOK)
	return nil
}

func NewTestController() *TestController {
	thisController := &controller.Controller{
		Name: "RootController",
	}
	return &TestController{
		Controller: *thisController,
	}
}
