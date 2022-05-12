package router_trie

import "github.com/bludot/gorouter/controller"

type IRouterTrie interface {
	Insert(key string, controller controller.IController)
	GetController(key string) (controlelr *controller.IController, params *map[string]string, found bool)
}
