package router_test

import (
	"context"
	controller2 "github.com/bludot/gorouter/core/controller"
	"github.com/bludot/gorouter/core/router"
	"github.com/bludot/gorouter/core/router/entities"
	"log"
	"testing"

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
	controller2.Controller
}

func (c *RootController) Handle(ctx context.Context, params *entities.RouteParams, queryParams *entities.QueryParams) error {
	log.Println("Controller:", c.Name)
	log.Println("Params:", params)
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
	t.Run("should process a sub route", func(t *testing.T) {
		a := assert.New(t)

		routerService := router.NewRouter()

		routerService.AddRoute(router.Route{
			Controller: NewRootController(),
			Path:       "/test",
		})

		err := routerService.Process(context.TODO(), "/test")

		a.NoError(err)

	})

	t.Run("should process a sub route with params", func(t *testing.T) {
		a := assert.New(t)

		routerService := router.NewRouter()

		routerService.AddRoute(router.Route{
			Controller: NewRootController(),
			Path:       "/test/$test",
		})

		err := routerService.Process(context.TODO(), "/test/123")

		a.NoError(err)

	})
}
