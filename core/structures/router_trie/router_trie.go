package router_trie

import (
	"errors"
	"github.com/bludot/gorouter/core/router/entities"
	"log"
	"strings"
)

var (
	NotFoundError = errors.New("Not found")
)

type Node struct {
	children map[string]*Node
	isWord   bool
	value    *entities.Route
	params   *[]string
}

func NewNode(route entities.Route, params *[]string) *Node {
	return &Node{
		children: make(map[string]*Node),
		value:    &route,
		params:   params,
	}
}

type RouterTrie struct {
	Root *Node
}

func NewRouteTrie() IRouterTrie {
	return &RouterTrie{
		Root: &Node{
			children: make(map[string]*Node),
		},
	}
}
func (t *RouterTrie) Insert(route entities.Route) {
	node := t.Root
	key := route.Path
	for part, i := PathSegmenter(key, 0); part != ""; part, i = PathSegmenter(key, i) {
		if len(part) > 1 && part[1] == '$' {
			params := node.params
			if params == nil {
				paramsMade := make([]string, 0)
				params = &paramsMade
			}
			*params = append(*params, part[2:])
			node.params = params

		} else {
			child, _ := node.children[part]
			if child == nil {
				if node.children == nil {
					node.children = map[string]*Node{}
				}
				child = NewNode(route, nil)
				node.children[part] = child
			}
			node = child
		}
	}

}

func PathSegmenter(path string, start int) (segment string, next int) {
	if len(path) == 0 || start < 0 || start > len(path)-1 {
		return "", -1
	}
	end := strings.IndexRune(path[start+1:], '/') // next '/' after 0th rune
	if end == -1 {
		return path[start:], -1
	}
	return path[start : start+end+1], start + end + 1
}

func (t *RouterTrie) GetController(key string) (route *entities.Route, params *entities.RouteParams, err error) {
	log.Println("GetController", key)
	foundNode, params, found := t.Search(key)
	var thisError error
	if !found {
		thisError = NotFoundError
		return nil, params, thisError
	}
	return foundNode.value, params, thisError
}

func (t *RouterTrie) Search(key string) (*Node, *entities.RouteParams, bool) {
	node := t.Root
	params := make(entities.RouteParams)
	for part, i := PathSegmenter(key, 0); part != ""; part, i = PathSegmenter(key, i) {
		child, _ := node.children[part]
		skip := 0
		if child == nil {
			return nil, &params, false
		}
		if child.params != nil {
			var newPart string
			end := i
			newPart, _ = PathSegmenter(key, end)

			if child.children[newPart] == nil && end != -1 {

				for _, param := range *child.params {
					newPart, end = PathSegmenter(key, end)
					if end == -1 {
						params[param] = newPart[1:]
						skip += len(newPart)
					} else {
						params[param] = newPart[1:]
						skip = end
					}
				}
				// split string by '/'
				newPart, end = PathSegmenter(key, skip)
				node = child
				if end == -1 {
					return node, &params, true
					// return node, &params, true
				}
			}

		}
		i += skip

		node = child
	}
	return node, &params, true
}
