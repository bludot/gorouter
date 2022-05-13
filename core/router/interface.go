package router

import (
	"context"
	"github.com/bludot/gorouter/core/router/entities"
	"net/http"
)

type Router interface {
	AddRoute(route Route)
	GetRoutes() []Route
	ParseQueryParams(path string) *entities.QueryParams
	Process(ctx context.Context, path string) error
	ServeHTTP(http.ResponseWriter, *http.Request)
}
