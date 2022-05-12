package controller

import (
	"context"
	"log"

	"github.com/bludot/gorouter/router/entities"
)

type Controller struct {
	Name string `json:"name"`
}

func (c *Controller) Run(ctx context.Context, params *entities.RouteParams, queryParams *entities.QueryParams) error {
	log.Println("Controller:", c.Name)
	log.Println("Params:", params)
	return nil
}
