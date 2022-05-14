package router_trie

import (
	"github.com/bludot/gorouter/core/router/entities"
)

type IRouterTrie interface {
	Insert(route entities.Route)
	GetController(key string) (handler *entities.Route, params *entities.RouteParams, err error)
}
