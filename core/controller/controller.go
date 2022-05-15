package controller

import (
	"context"
	"github.com/bludot/gorouter/core/router/entities"
	"github.com/bludot/gorouter/core/service"
	"log"
)

type Controller struct {
	Name     string            `json:"name"`
	Services []service.Service `json:"services"`
}

func (c *Controller) Handle(ctx context.Context, r entities.HTTPRequest) error {
	log.Println("Handler:", c.Name)
	log.Println("Params:", r.Params)
	return nil
}

func (c *Controller) GetService(name string) *service.Service {
	for _, s := range c.Services {
		if s.GetName() == name {
			return &s
		}
	}
	return nil
}
