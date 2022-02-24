package tree

import (
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	trie := NewTrie()

	trie.Insert("apple")
	trie.Insert("b")
	trie.Insert("abc")
	trie.Insert("dddddddd")

	ret := trie.Search("abc")
	if !ret {
		t.Errorf("wrong result, expected true but got %t", ret)
	}

	ret = trie.Search("abandon")
	if ret {
		t.Errorf("wrong result, expected false but got %t", ret)
	}

	trie.Delete("abc")
	ret = trie.Search("abc")
	if ret {
		t.Errorf("wrong result, expected false but got %t", ret)
	}
}

func TestTrie_StartsWith(t *testing.T) {
	trie := NewTrie()

	trie.Insert("a")
	trie.Insert("b")
	trie.Insert("c")
	trie.Insert("d")
	trie.Insert("e")
	trie.Insert("f")
	trie.Insert("g")
	trie.Insert("h")
	trie.Insert("i")
	trie.Insert("j")
	trie.Insert("k")
	trie.Insert("l")
	trie.Insert("m")
	trie.Insert("n")
	trie.Insert("o")
	trie.Insert("p")
	trie.Insert("q")
	trie.Insert("r")
	trie.Insert("s")
	trie.Insert("t")
	trie.Insert("u")
	trie.Insert("v")
	trie.Insert("w")
	trie.Insert("x")
	trie.Insert("y")
	trie.Insert("z")

	ret := trie.StartsWith("aa")
	if ret {
		t.Errorf("wrong result, expected false but got %t", ret)
	}
}
