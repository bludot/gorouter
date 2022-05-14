package entities

import "context"

type QueryParams map[string]string

type RouteParams map[string]string

func (r RouteParams) Get(key string) string {
	return r[key]
}

type RouteHandler = func(ctx context.Context, params *RouteParams, queryParams *QueryParams) error

type Route struct {
	Handler RouteHandler
	Path    string
	Method  string
}
