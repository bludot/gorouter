package trie

import (
	"github.com/bludot/gorouter/controller"
	"strings"
)

type Node struct {
	children   map[string]*Node
	isWord     bool
	controller controller.IController
}

type Trie struct {
	Root *Node
}

func NewTrie() ITrie {
	return &Trie{
		Root: &Node{
			children: make(map[string]*Node),
		},
	}
}

func NewNode(controller controller.IController) *Node {
	return &Node{
		children:   make(map[string]*Node),
		controller: controller,
	}
}

func (t *Trie) Insert(key string, controller controller.IController) {
	node := t.Root
	for part, i := PathSegmenter(key, 0); part != ""; part, i = PathSegmenter(key, i) {
		child, _ := node.children[part]
		if child == nil {
			if node.children == nil {
				node.children = map[string]*Node{}
			}
			child = NewNode(controller)
			node.children[part] = child
		}
		node = child
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

func (t *Trie) Search(key string) (controller.IController, bool) {
	node := t.Root
	for part, i := PathSegmenter(key, 0); part != ""; part, i = PathSegmenter(key, i) {
		child, _ := node.children[part]
		if child == nil {
			return nil, false
		}
		node = child
	}
	return node.controller, true
}

func (t *Trie) Delete(key string) {
	node := t.Root
	for part, i := PathSegmenter(key, 0); part != ""; part, i = PathSegmenter(key, i) {
		child, _ := node.children[part]
		if child == nil {
			return
		}
		node = child
	}
	node.isWord = false
}
