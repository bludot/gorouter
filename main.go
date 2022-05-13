package main

import (
	"context"
	controller2 "github.com/bludot/gorouter/core/controller"
	"github.com/bludot/gorouter/core/renderer"
	"github.com/bludot/gorouter/core/router"
	"github.com/bludot/gorouter/core/router/entities"
	"github.com/bludot/gorouter/core/template"
	"github.com/bludot/gorouter/core/transformer"
	"log"
	"net/http"
)

type RootController struct {
	controller2.Controller
}

func (c *RootController) Handle(ctx context.Context, params *entities.RouteParams, queryParams *entities.QueryParams) error {
	log.Println("Controller:", c.Name)
	log.Println("Params:", params)
	renderer.GetRender().Render("index.html", map[string]string{
		"body": "hello world",
	})
	return nil
}

func NewRootController() controller2.IController {
	thisController := &controller2.Controller{
		Name: "RootController",
	}
	return &RootController{
		Controller: *thisController,
	}
}

type TestController struct {
	controller2.Controller
}

func (c *TestController) Handle(ctx context.Context, params *entities.RouteParams, queryParams *entities.QueryParams) error {
	log.Println("Controller:", c.Name)
	log.Println("Params:", params)
	renderer.GetRender().Render("index.html", map[string]string{
		"body": "bye world",
	})
	return nil
}

func NewTestController() controller2.IController {
	thisController := &controller2.Controller{
		Name: "RootController",
	}
	return &TestController{
		Controller: *thisController,
	}
}

func main() {
	renderer.
		GetRender().
		SetTemplateEngine(template.GetTemplateEngine()).
		SetTransformer(transformer.GetTransformer())
	mainRouter := router.NewRouter()
	mainRouter.AddRoute(router.Route{
		Controller: NewRootController(),
		Path:       "/",
	})
	mainRouter.AddRoute(router.Route{
		Controller: NewRootController(),
		Path:       "/test/$id",
	})
	mainRouter.AddRoute(router.Route{
		Controller: NewTestController(),
		Path:       "/test",
	})
	http.Handle("/", mainRouter)
	http.ListenAndServe(":8080", nil)
}
