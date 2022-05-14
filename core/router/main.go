package router

import (
	"context"
	"errors"
	"github.com/bludot/gorouter/core/renderer"
	"github.com/bludot/gorouter/core/router/entities"
	"github.com/bludot/gorouter/core/structures/router_trie"
	"log"
	"net/http"
	"strings"
)

type RouterService struct {
	routes []entities.Route
	cached map[string]*entities.Route
	Trie   router_trie.IRouterTrie
}

func NewRouter() Router {
	newTrie := router_trie.NewRouteTrie()
	return &RouterService{
		routes: make([]entities.Route, 0),
		cached: make(map[string]*entities.Route),
		Trie:   newTrie,
	}
}

func (r *RouterService) AddRoute(route entities.Route) {
	r.routes = append(r.routes, route)
	r.cached[route.Path] = &route
	r.Trie.Insert(route)
}

func (r *RouterService) GetRoutes() []entities.Route {
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

func (r *RouterService) Process(ctx context.Context, req *http.Request, path string) error {

	routePath := strings.Split(path, "?")
	var queryParams *entities.QueryParams
	if len(routePath) > 1 {
		queryParams = r.ParseQueryParams(routePath[1])
	}
	route, params, err := r.Trie.GetController(routePath[0])
	if err != nil {
		log.Println("error is not nil", err)
		return err
	}
	if req.Method != route.Method {
		renderer.GetRender().ToJSON(map[string]interface{}{"error": "Method not allowed"}, http.StatusInternalServerError)
		return errors.New("method not allowed")
	}

	log.Println("error is nil", err)
	return (*route).Handler(ctx, params, queryParams)

}

func (r *RouterService) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	log.Println("Request: ", path)
	err := r.Process(context.Background(), req, path)
	if err != nil {
		log.Println("Error: ", err)
	}
	renderer.GetRender().Respond(w, req)
}
