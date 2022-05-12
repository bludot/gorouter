package router

import (
	"context"
	"net/http"

	"github.com/bludot/gorouter/router/entities"
)

type Router interface {
	AddRoute(route Route)
	GetRoutes() []Route
	ParseQueryParams(path string) *entities.QueryParams
	Process(ctx context.Context, path string) error
	ServeHTTP(http.ResponseWriter, *http.Request)
}
