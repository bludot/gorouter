package entities

import (
	"context"
	"net/http"
)

type QueryParams map[string]string

type RouteParams map[string]string

func (r RouteParams) Get(key string) string {
	return r[key]
}

type HTTPRequest struct {
	Header     http.Header
	Method     string
	URL        string
	Body       *[]byte
	RemoteAddr string
	Query      *QueryParams
	Params     *RouteParams
}

type RouteHandler = func(ctx context.Context, request HTTPRequest) error

type Route struct {
	Handler RouteHandler
	Path    string
	Method  string
}
