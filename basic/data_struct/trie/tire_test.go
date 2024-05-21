package trie_test

import (
	"fmt"
	"gostudy/basic/data_struct/trie"
	"testing"
)

func TestTire(t *testing.T) {
	node := trie.New()
	node.Insert("hello")
	node.Insert("world")
	node.Insert("hel")
	fmt.Println(node.Search("hello"))
	fmt.Println(node.Search("h"))
	fmt.Println(node.Search("w"))
	fmt.Println(node.Search("e"))
}
