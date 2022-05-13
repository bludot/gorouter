package router_trie

import (
	"github.com/bludot/gorouter/core/controller"
	"github.com/bludot/gorouter/core/router/entities"
)

type IRouterTrie interface {
	Insert(key string, controller controller.IController)
	GetController(key string) (controller *controller.IController, params *entities.RouteParams, err error)
}
