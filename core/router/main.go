package router

import (
	"context"
	"errors"
	"github.com/bludot/gorouter/core/renderer"
	"github.com/bludot/gorouter/core/router/entities"
	"github.com/bludot/gorouter/core/structures/router_trie"
	"io"
	"log"
	"net/http"
	"strings"
)

type RouterService struct {
	routes []entities.Route
	cached map[string]*entities.Route
	Tries  map[string]router_trie.IRouterTrie
}

func NewRouter() Router {

	return &RouterService{
		routes: make([]entities.Route, 0),
		cached: make(map[string]*entities.Route),
		Tries:  map[string]router_trie.IRouterTrie{},
	}
}

func (r *RouterService) AddRoute(route entities.Route) {
	r.routes = append(r.routes, route)
	r.cached[route.Path] = &route
	if _, ok := r.Tries[route.Method]; !ok {
		r.Tries[route.Method] = router_trie.NewRouteTrie()
	}
	r.Tries[route.Method].Insert(route)
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

func (r *RouterService) Process(ctx context.Context, req *http.Request) error {
	path := req.URL.Path
	routePath := strings.Split(path, "?")
	var queryParams *entities.QueryParams
	if len(routePath) > 1 {
		queryParams = r.ParseQueryParams(routePath[1])
	}
	if _, ok := r.Tries[req.Method]; !ok {
		renderer.GetRender().ToJSON(map[string]interface{}{"error": "Method not allowed"}, http.StatusMethodNotAllowed)
		return errors.New("Method not found")
	}
	route, params, err := r.Tries[req.Method].GetController(routePath[0])
	if err != nil {
		renderer.GetRender().ToJSON(map[string]interface{}{"error": "Method not allowed"}, http.StatusMethodNotAllowed)
		log.Println("error is not nil", err)
		return errors.New("method not allowed")
	}

	log.Println("error is nil", err)
	b, err := io.ReadAll(req.Body)
	return (*route).Handler(ctx, entities.HTTPRequest{
		Body:       &b,
		Params:     params,
		Query:      queryParams,
		Header:     req.Header,
		URL:        req.URL.String(),
		Method:     req.Method,
		RemoteAddr: req.RemoteAddr,
	})

}

func (r *RouterService) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := r.Process(context.Background(), req)
	if err != nil {
		log.Println("Error: ", err)
	}
	renderer.GetRender().Respond(w, req)
}
