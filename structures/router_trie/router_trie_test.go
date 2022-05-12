package router_trie_test

import (
	"context"
	"log"
	"testing"

	"github.com/bludot/gorouter/controller"
	"github.com/bludot/gorouter/router/entities"
	"github.com/bludot/gorouter/structures/router_trie"
	"github.com/stretchr/testify/assert"
)

type AController struct {
	controller.Controller
}

func (c *AController) Run(ctx context.Context, params *entities.RouteParams, queryParams *entities.QueryParams) error {
	log.Println("Controller:", c.Name)
	log.Println("Params:", params)
	return nil
}

func NewAController() controller.IController {
	thisController := &controller.Controller{
		Name: "AController",
	}
	return &AController{
		Controller: *thisController,
	}
}

type RootController struct {
	controller.Controller
}

func (c *RootController) Run(ctx context.Context, params *entities.RouteParams, queryParams *entities.QueryParams) error {
	log.Println("Controller:", c.Name)
	log.Println("Params:", params)
	return nil
}

func NewRootController() controller.IController {
	thisController := &controller.Controller{
		Name: "RootController",
	}
	return &RootController{
		Controller: *thisController,
	}
}

func TestRouterTrie(t *testing.T) {
	a := assert.New(t)

	node := router_trie.NewRouteTrie()
	node.Insert("/", NewRootController())
	node.Insert("/a", NewAController())
	node.Insert("/b", NewAController())
	node.Insert("/a/b", NewRootController())
	node.Insert("/a/b/c", NewRootController())
	node.Insert("/a/b/$test/$test2", NewRootController())
	node.Insert("/a/b/d/$test2", NewRootController())
	control, params, err := node.GetController("/a")
	a.NotNil(control)
	a.Equal(len(*params), 0)
	a.NoError(err)

	control, params, err = node.GetController("/")
	a.NotNil(control)
	a.Equal(len(*params), 0)
	a.NoError(err)

	control, params, err = node.GetController("/a/b/123/1234")
	a.NotNil(control)
	a.Equal(len(*params), 2)
	a.NoError(err)

}
