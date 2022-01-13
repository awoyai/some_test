package trietree

import (
	"fmt"
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	trie := Constructor()
	trie.Insert("abc")
	trie.Insert("acd")
	trie.Insert("acdd")
	b := trie.Search("acd")
	b2 := trie.Search("ac")
	b3 := trie.StartsWith("ac")
	println(b, b2, b3)
	fmt.Printf("%+v", trie)
}

func TestTrie_Inserts(t *testing.T) {
	trie := ConstructorM()
	trie.Insert("abc")
	trie.Insert("acd")
	trie.Insert("acdd")
	b := trie.Search("acd")
	b2 := trie.Search("ac")
	b3 := trie.StartsWith("ac")
	println(b, b2, b3)
	fmt.Printf("%+v", trie)
}