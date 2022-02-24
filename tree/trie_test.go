package tree

import (
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	trie := NewTrie()

	trie.Insert("a")
	trie.Insert("ab")
	trie.Insert("abc")
	trie.Insert("aca")
	trie.Insert("acb")
	trie.Insert("acbc")

	ret := trie.Search("abc")
	if !ret {
		t.Errorf("wrong result, expected true but got %t", ret)
	}

	ret = trie.Search("abd")
	if ret {
		t.Errorf("wrong result, expected false but got %t", ret)
	}

	trie.Delete("abc")
	ret = trie.Search("abc")
	if ret {
		t.Errorf("wrong result, expected false but got %t", ret)
	}
}
