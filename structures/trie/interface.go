package trie

import "github.com/bludot/gorouter/controller"

type ITrie interface {
	Insert(key string, controller controller.IController)
	Search(key string) (controller.IController, bool)
	/*Keys() []string
	Values() []interface{}
	PrefixSearch(key string) []string
	PrefixSearchValues(key string) []interface{}
	PrefixSearchKeys(key string) []string*/
}
