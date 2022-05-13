package controller

import (
	"context"
	"github.com/bludot/gorouter/core/router/entities"
)

type IController interface {
	Handle(ctx context.Context, params *entities.RouteParams, queryParams *entities.QueryParams) error
}
