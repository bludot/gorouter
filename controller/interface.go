package controller

import (
	"context"

	"github.com/bludot/gorouter/router/entities"
)

type IController interface {
	Run(ctx context.Context, params *entities.RouteParams, queryParams *entities.QueryParams) error
}
