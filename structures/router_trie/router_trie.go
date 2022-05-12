package router_trie

import (
	"github.com/bludot/gorouter/controller"
	"strings"
)

type Node struct {
	children map[string]*Node
	isWord   bool
	value    *controller.IController
	params   *[]string
}

func NewNode(controller controller.IController, params *[]string) *Node {
	return &Node{
		children: make(map[string]*Node),
		value:    &controller,
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
func (t *RouterTrie) Insert(key string, controller controller.IController) {
	node := t.Root

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
				child = NewNode(controller, nil)
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

func (t *RouterTrie) GetController(key string) (controller *controller.IController, params *map[string]string, found bool) {
	foundNode, params, found := t.Search(key)

	return foundNode.value, params, found
}

func (t *RouterTrie) Search(key string) (*Node, *map[string]string, bool) {
	node := t.Root
	params := make(map[string]string)
	for part, i := PathSegmenter(key, 0); part != ""; part, i = PathSegmenter(key, i) {
		child, _ := node.children[part]
		skip := 0
		if child == nil {
			return nil, &params, false
		}
		if child.params != nil {
			var newPart string
			for j, param := range *child.params {
				newPart, _ = PathSegmenter(key, i+j)
				params[param] = newPart
				skip = j
			}
			// split string by '/'
			split := strings.Split(key, "/")
			if len(split[i+skip:]) == 0 {
				node = child
				return node, &params, true
			}

		}
		i += skip

		node = child
	}
	return node, &params, true
}
