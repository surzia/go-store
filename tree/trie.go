package tree

import "sync"

type TrieNode struct {
	children map[string]*TrieNode
	value    string
	end      bool
}

func NewTrieNode(e string) *TrieNode {
	ret := make(map[string]*TrieNode)
	return &TrieNode{
		children: ret,
		value:    e,
		end:      true,
	}
}

type Trie struct {
	sync.RWMutex
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[string]*TrieNode),
		},
	}
}

// Insert a node into trie, the complexity of this operation is O(n),
// where n represents the key size.
func (t *Trie) Insert(word string) {
	t.Lock()
	defer t.Unlock()

	current := t.root

	for i := 0; i < len(word); i++ {
		if child, ok := current.children[string(word[i])]; ok {
			current = child
		} else {
			current.children[string(word[i])] = NewTrieNode(string(word[i]))
			current = current.children[string(word[i])]
		}
	}

	current.end = true
}

// Search a node in trie and return true or false, the complexity
// of this algorithm is O(n), where n represents the length of the key.
func (t *Trie) Search(word string) bool {
	t.Lock()
	defer t.Unlock()

	current := t.root

	for i := 0; i < len(word); i++ {
		node := current.children[string(word[i])]
		if node == nil {
			return false
		}
		current = node
	}
	return current.end
}

// Delete a node from trie, the complexity of this algorithm is O(n),
// where n represents the length of the key.
func (t *Trie) Delete(word string) {
	t.Lock()
	defer t.Unlock()

	_ = deleteNode(t.root, word, 0)
}

// deleteNode a node from specific node in trie
func deleteNode(cur *TrieNode, word string, index int) bool {
	if index == len(word) {
		if !cur.end {
			return false
		}
		cur.end = false
		return len(cur.children) == 0
	}

	node := cur.children[string(word[index])]
	if node == nil {
		return false
	}
	shouldDeleteCurrentNode := deleteNode(node, word, index+1) && !node.end

	if shouldDeleteCurrentNode {
		delete(cur.children, string(word[index]))
		return len(cur.children) == 0
	}

	return false
}

// StartsWith returns true if there is a previously inserted string word
// that has the prefix, and false otherwise.
func (t *Trie) StartsWith(prefix string) bool {
	current := t.root

	for i := 0; i < len(prefix); i++ {
		node, ok := current.children[string(prefix[i])]
		if !ok {
			return false
		}
		current = node
	}
	return true
}
