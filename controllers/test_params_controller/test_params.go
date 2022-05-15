package test_params_controller

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

func (c *TestParamsController) TestParams(ctx context.Context, r entities.HTTPRequest) error {
	log.Println("Handler:", c.Name)
	log.Println("Params:", r.Params)
	renderer.GetRender().Render("index.html", map[string]string{
		"body": "this is a test (POST)",
	}, http.StatusOK)
	return nil
}

func (c *TestParamsController) GetParams(ctx context.Context, r entities.HTTPRequest) error {
	renderer.GetRender().Render("index.html", map[string]string{
		"body": "this is a test (GET)",
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
