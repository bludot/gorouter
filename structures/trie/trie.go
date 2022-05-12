package trie

import (
	"log"
	"strings"
)

type Node struct {
	Children map[string]*Node
	isWord   bool
	Value    interface{}
}

type Trie struct {
	Root *Node
}

func NewTrie() ITrie {
	return &Trie{
		Root: &Node{
			Children: make(map[string]*Node),
		},
	}
}

func NewNode(value interface{}, params *[]string) *Node {
	return &Node{
		Children: make(map[string]*Node),
		Value:    value,
	}
}

func (t *Trie) Insert(key string, value interface{}) {
	node := t.Root

	for part, i := PathSegmenter(key, 0); part != ""; part, i = PathSegmenter(key, i) {
		child, _ := node.Children[part]
		if child == nil {
			if node.Children == nil {
				node.Children = map[string]*Node{}
			}
			child = NewNode(value, nil)
			node.Children[part] = child
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

func (t *Trie) Search(key string) (*Node, bool) {
	node := t.Root
	for part, i := PathSegmenter(key, 0); part != ""; part, i = PathSegmenter(key, i) {
		log.Println(part)
		child, _ := node.Children[part]
		lookforwardPart, j := PathSegmenter(key, i)
		log.Println("forward", lookforwardPart, j)
		if child == nil {
			return nil, false
		}
		node = child
	}
	return node, true
}

func (t *Trie) Delete(key string) {
	node := t.Root
	for part, i := PathSegmenter(key, 0); part != ""; part, i = PathSegmenter(key, i) {
		child, _ := node.Children[part]
		if child == nil {
			return
		}
		node = child
	}
	node.isWord = false
}
