package trie_test

import (
	"github.com/bludot/gorouter/controller"
	"github.com/bludot/gorouter/structures/trie"
	"github.com/stretchr/testify/assert"
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

	a := assert.New(t)
	node := trie.NewTrie()
	node.Insert("/", "tset")
	node.Insert("/a", "tset2")
	node.Insert("/b", "test3")
	node.Insert("/a/b", "test4")
	node.Insert("/a/b/c", "test5")
	node.Insert("/a/b/$test/$test2", "test6")
	node.Insert("/a/b/d/$test2", "test7")
	log.Println(node.Search("/a"))
	val, _ := node.Search("/a")
	a.Equal("tset2", val)

	val, _ = node.Search("/a/b")
	a.Equal("test4", val)

	val, _ = node.Search("/a/b/c")
	a.Equal("test5", val)

	val, _ = node.Search("/a/b/$test/$test2")
	a.Equal("test6", val)

	val, _ = node.Search("/a/b/d/$test2")
	a.Equal("test7", val)

	val, _ = node.Search("/a/b/d/$test2")
	a.Equal("test7", val)

}
