package router_trie_test

import (
	"github.com/bludot/gorouter/controller"
	"github.com/bludot/gorouter/structures/router_trie"

	"log"
	"testing"
)

type AController struct {
	controller.Controller
}

func (c *AController) Run(params map[string]string) (string, error) {
	log.Println("Controller:", c.Name)
	log.Println("Params:", params)
	return "", nil
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

func (c *RootController) Run(params map[string]string) (string, error) {
	log.Println("Controller:", c.Name)
	log.Println("Params:", params)
	return "", nil
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

	node := router_trie.NewRouteTrie()
	node.Insert("/", NewRootController())
	node.Insert("/a", NewAController())
	node.Insert("/b", NewAController())
	node.Insert("/a/b", NewRootController())
	node.Insert("/a/b/c", NewRootController())
	node.Insert("/a/b/$test/$test2", NewRootController())
	node.Insert("/a/b/d/$test2", NewRootController())
	log.Println(node.GetController("/a"))
	control, _, _ := node.GetController("/a")
	(*control).Run(map[string]string{})

	control, _, _ = node.GetController("/")
	(*control).Run(map[string]string{})

	control, params, exists := node.GetController("/a/b/123/123")
	if exists {
		log.Println(params)
		(*control).Run(*params)
	}

}
