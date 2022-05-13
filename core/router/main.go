package router

import (
	"context"
	"github.com/bludot/gorouter/core/controller"
	"github.com/bludot/gorouter/core/renderer"
	"github.com/bludot/gorouter/core/router/entities"
	router_trie2 "github.com/bludot/gorouter/core/structures/router_trie"
	"log"
	"net/http"
	"strings"
)

type Route struct {
	Controller controller.IController
	Path       string
}

type RouterService struct {
	routes []Route
	cached map[string]*Route
	Trie   router_trie2.IRouterTrie
}

func NewRouter() Router {
	newTrie := router_trie2.NewRouteTrie()
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
	if err == nil {
		log.Println("error is nil", err)
		return (*controller).Handle(ctx, params, queryParams)
	}
	log.Println("error is not nil", err)
	return err
}

func (r *RouterService) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	log.Println("Request: ", path)
	err := r.Process(context.Background(), path)
	if err != nil {
		log.Println("Error: ", err)
	}
	renderer.GetRender().Respond(w, req)
}
