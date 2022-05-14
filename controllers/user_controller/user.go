package user_controller

import (
	"context"
	"github.com/bludot/gorouter/core/controller"
	"github.com/bludot/gorouter/core/renderer"
	"github.com/bludot/gorouter/core/router/entities"
	"log"
	"net/http"
)

package root

import (
"context"
controller "github.com/bludot/gorouter/core/controller"
"github.com/bludot/gorouter/core/renderer"
"github.com/bludot/gorouter/core/router/entities"
"log"
"net/http"
)

type UserController struct {
	controller.Controller
}

func (c *UserController) GetUser(ctx context.Context, params *entities.RouteParams, queryParams *entities.QueryParams) error {
	log.Println("Handler:", c.Name)
	log.Println("Params:", params)
	renderer.GetRender().ToJSON(map[string]string{"hello": "world"}, http.StatusOK)
	return nil
}

func Controller() *UserController {
	thisController := &controller.Controller{
		Name: "UserController",
	}
	return &UserController{
		Controller: *thisController,
	}
}

