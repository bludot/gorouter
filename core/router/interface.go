package router

import (
	"context"
	"github.com/bludot/gorouter/core/router/entities"
	"net/http"
)

type Router interface {
	AddRoute(route entities.Route)
	GetRoutes() []entities.Route
	ParseQueryParams(path string) *entities.QueryParams
	Process(ctx context.Context, req *http.Request, path string) error
	ServeHTTP(http.ResponseWriter, *http.Request)
}
