package router

import (
	"context"
	"strings"

	"github.com/bludot/gorouter/controller"
	"github.com/bludot/gorouter/router/entities"
	"github.com/bludot/gorouter/structures/router_trie"
)

type Route struct {
	Controller controller.IController
	Path       string
}

type RouterService struct {
	routes []Route
	cached map[string]*Route
	Trie   router_trie.IRouterTrie
}

func NewRouter() Router {
	newTrie := router_trie.NewRouteTrie()
	return &RouterService{
		routes: make([]Route, 0),
		cached: make(map[string]*Route),
		Trie:   newTrie,
	}
}

func (r *RouterService) AddRoute(route Route) {
	r.routes = append(r.routes, route)
	r.cached[route.Path] = &route
	r.Trie.Insert(route.Path, route.Controller)
}

func (r *RouterService) GetRoutes() []Route {
	return r.routes
}

func (r *RouterService) ParseQueryParams(path string) *entities.QueryParams {
	queryParams := make(entities.QueryParams)
	params := strings.Split(path, "&")
	for _, param := range params {
		keyValue := strings.Split(param, "=")
		queryParams[keyValue[0]] = keyValue[1]
	}
	return &queryParams
}

func (r *RouterService) Process(ctx context.Context, path string) error {
	route := strings.Split(path, "?")
	var queryParams *entities.QueryParams
	if len(route) > 1 {
		queryParams = r.ParseQueryParams(route[1])
	}
	controller, params, err := r.Trie.GetController(route[0])
	if err != nil {
		return (*controller).Run(ctx, params, queryParams)
	}
	return err
}
