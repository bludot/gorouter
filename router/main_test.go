package router_test

import (
	"context"
	"log"
	"testing"

	"github.com/bludot/gorouter/controller"
	router "github.com/bludot/gorouter/router"
	"github.com/bludot/gorouter/router/entities"
	"github.com/stretchr/testify/assert"
)

func TestNewRouter(t *testing.T) {
	t.Run("should return a new router", func(t *testing.T) {
		routerService := router.NewRouter()
		if routerService == nil {
			t.Errorf("Expected a router, got nil")
		}
	})
}

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

func TestRouterService_AddRoute(t *testing.T) {

	t.Run("should add a route", func(t *testing.T) {
		routerService := router.NewRouter()

		routerService.AddRoute(router.Route{
			Controller: NewRootController(),
			Path:       "/",
		})
	})
}

func TestRouterService_Process(t *testing.T) {
	t.Run("should process a route", func(t *testing.T) {
		a := assert.New(t)
		routerService := router.NewRouter()

		routerService.AddRoute(router.Route{
			Controller: NewRootController(),
			Path:       "/",
		})

		err := routerService.Process(context.TODO(), "/")

		a.NoError(err)

	})
	t.Run("should not process a route with queryParams", func(t *testing.T) {
		a := assert.New(t)

		routerService := router.NewRouter()

		routerService.AddRoute(router.Route{
			Controller: NewRootController(),
			Path:       "/",
		})

		err := routerService.Process(context.TODO(), "/?test=test")

		a.NoError(err)

	})
}
