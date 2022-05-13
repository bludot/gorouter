package trie_test

import (
	"github.com/bludot/gorouter/core/structures/trie"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	a.Equal("tset2", val.Value.(string))

	val, _ = node.Search("/a/b")
	a.Equal("test4", val.Value.(string))

	val, _ = node.Search("/a/b/c")
	a.Equal("test5", val.Value.(string))

	val, _ = node.Search("/a/b/$test/$test2")
	a.Equal("test6", val.Value.(string))

	val, _ = node.Search("/a/b/d/$test2")
	a.Equal("test7", val.Value.(string))

	val, _ = node.Search("/a/b/d/$test2")
	a.Equal("test7", val.Value.(string))

}
