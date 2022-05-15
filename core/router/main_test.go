package router_test

import (
	"context"
	controller2 "github.com/bludot/gorouter/core/controller"
	"github.com/bludot/gorouter/core/router"
	"github.com/bludot/gorouter/core/router/entities"
	"log"
	"net/http"
	"net/url"
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

func (c *RootController) Handle(ctx context.Context, r entities.HTTPRequest) error {
	log.Println("Handler:", c.Name)
	log.Println("Params:", r.Params)
	return nil
}

func NewRootController() *RootController {
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

		routerService.AddRoute(entities.Route{
			Handler: NewRootController().Handle,
			Path:    "/",
			Method:  "GET",
		})
	})
}

func TestRouterService_Process(t *testing.T) {
	t.Run("should process a route", func(t *testing.T) {
		a := assert.New(t)
		routerService := router.NewRouter()

		routerService.AddRoute(entities.Route{
			Handler: NewRootController().Handle,
			Path:    "/",
			Method:  "GET",
		})

		req := http.Request{
			Method: "GET",
			URL: &url.URL{
				Scheme:      "",
				Opaque:      "",
				User:        nil,
				Host:        "",
				Path:        "/",
				RawPath:     "",
				ForceQuery:  false,
				RawQuery:    "",
				Fragment:    "",
				RawFragment: "",
			},
		}

		err := routerService.Process(context.TODO(), &req)

		a.NoError(err)

	})
	t.Run("should not process a route with queryParams", func(t *testing.T) {
		a := assert.New(t)

		routerService := router.NewRouter()

		routerService.AddRoute(entities.Route{
			Handler: NewRootController().Handle,
			Path:    "/",
			Method:  "GET",
		})

		req := http.Request{
			Method: "GET",
			URL: &url.URL{
				Scheme:      "",
				Opaque:      "",
				User:        nil,
				Host:        "",
				Path:        "/?test=test",
				RawPath:     "",
				ForceQuery:  false,
				RawQuery:    "",
				Fragment:    "",
				RawFragment: "",
			},
		}

		err := routerService.Process(context.TODO(), &req)

		a.NoError(err)

	})
	t.Run("should process a sub route", func(t *testing.T) {
		a := assert.New(t)

		routerService := router.NewRouter()

		routerService.AddRoute(entities.Route{
			Handler: NewRootController().Handle,
			Path:    "/test",
			Method:  "GET",
		})

		req := http.Request{
			Method: "GET",
			URL: &url.URL{
				Scheme:      "",
				Opaque:      "",
				User:        nil,
				Host:        "",
				Path:        "/test",
				RawPath:     "",
				ForceQuery:  false,
				RawQuery:    "",
				Fragment:    "",
				RawFragment: "",
			},
		}

		err := routerService.Process(context.TODO(), &req)

		a.NoError(err)

	})

	t.Run("should process a sub route with params", func(t *testing.T) {
		a := assert.New(t)

		routerService := router.NewRouter()

		routerService.AddRoute(entities.Route{
			Handler: NewRootController().Handle,
			Path:    "/test/$test",
			Method:  "GET",
		})

		req := http.Request{
			Method: "GET",
			URL: &url.URL{
				Scheme:      "",
				Opaque:      "",
				User:        nil,
				Host:        "",
				Path:        "/test/123",
				RawPath:     "",
				ForceQuery:  false,
				RawQuery:    "",
				Fragment:    "",
				RawFragment: "",
			},
		}

		err := routerService.Process(context.TODO(), &req)

		a.NoError(err)

	})
}
