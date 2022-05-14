package router_trie_test

import (
	"context"
	controller2 "github.com/bludot/gorouter/core/controller"
	"github.com/bludot/gorouter/core/router/entities"
	"github.com/bludot/gorouter/core/structures/router_trie"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

type AController struct {
	controller2.Controller
}

func (c *AController) Handle(ctx context.Context, params *entities.RouteParams, queryParams *entities.QueryParams) error {
	log.Println("Handler:", c.Name)
	log.Println("Params:", params)
	return nil
}

func NewAController() controller2.IController {
	thisController := &controller2.Controller{
		Name: "AController",
	}
	return &AController{
		Controller: *thisController,
	}
}

type RootController struct {
	controller2.Controller
}

func (c *RootController) Handle(ctx context.Context, params *entities.RouteParams, queryParams *entities.QueryParams) error {
	log.Println("Handler:", c.Name)
	log.Println("Params:", params)
	return nil
}

func NewRootController() controller2.IController {
	thisController := &controller2.Controller{
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
	a.Equal(0, len(*params))
	a.NoError(err)

	control, params, err = node.GetController("/")
	a.NotNil(control)
	a.Equal(0, len(*params))
	a.NoError(err)

	control, params, err = node.GetController("/a/b/123/1234")
	a.NotNil(control)
	a.Equal(2, len(*params))
	a.NoError(err)

	control, params, err = node.GetController("/a/b/d/123")
	a.NotNil(control)
	log.Println("Params:", params)
	a.Equal(1, len(*params))
	a.NoError(err)

}
