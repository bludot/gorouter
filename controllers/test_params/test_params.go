package test_params

import (
	"context"
	controller "github.com/bludot/gorouter/core/controller"
	"github.com/bludot/gorouter/core/renderer"
	"github.com/bludot/gorouter/core/router/entities"
	"log"
	"net/http"
)

type TestParamsController struct {
	controller.Controller
}

func (c *TestParamsController) TestParams(ctx context.Context, params *entities.RouteParams, queryParams *entities.QueryParams) error {
	log.Println("Handler:", c.Name)
	log.Println("Params:", params)
	renderer.GetRender().Render("index.html", map[string]string{
		"body": "this is a test",
	}, http.StatusOK)
	return nil
}

func NewTestParamsController() *TestParamsController {
	thisController := &controller.Controller{
		Name: "RootController",
	}
	return &TestParamsController{
		Controller: *thisController,
	}
}
