package controller

import (
	"context"
	"github.com/bludot/gorouter/core/router/entities"
	"log"
)

type Controller struct {
	Name string `json:"name"`
}

func (c *Controller) Handle(ctx context.Context, params *entities.RouteParams, queryParams *entities.QueryParams) error {
	log.Println("Handler:", c.Name)
	log.Println("Params:", params)
	return nil
}
