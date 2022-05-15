package root_controller

import (
	"context"
	controller "github.com/bludot/gorouter/core/controller"
	"github.com/bludot/gorouter/core/renderer"
	"github.com/bludot/gorouter/core/router/entities"
	"log"
	"net/http"
)

type RootController struct {
	controller.Controller
}

func (c *RootController) Root(ctx context.Context, r entities.HTTPRequest) error {
	log.Println("Handler:", c.Name)
	log.Println("Params:", r.Params)
	renderer.GetRender().ToJSON(map[string]string{"hello": "world"}, http.StatusOK)
	return nil
}

func Controller() *RootController {
	thisController := &controller.Controller{
		Name: "RootController",
	}
	return &RootController{
		Controller: *thisController,
	}
}
