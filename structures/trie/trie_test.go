package trie_test

import (
	"github.com/bludot/gorouter/controller"
	"github.com/bludot/gorouter/structures/trie"
	"log"
	"testing"
)

type AController struct {
	controller.Controller
}

func (c *AController) Run(params []string) {
	log.Println("AController.Run", params)
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

func (c *RootController) Run(params []string) {
	log.Println("RootController.Run", params)
}

func NewRootController() controller.IController {
	thisController := &controller.Controller{
		Name: "RootController",
	}
	return &RootController{
		Controller: *thisController,
	}
}

func TestTrie(t *testing.T) {

	node := trie.NewTrie()
	node.Insert("/", NewRootController())
	node.Insert("/a", NewAController())
	log.Println(node.Search("/a"))
	control, _ := node.Search("/a")
	control.Run([]string{"tset"})

	control, _ = node.Search("/")
	control.Run([]string{"tset2"})
}
