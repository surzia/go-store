# README

Analysis and implementation of data structure in Go

## Usage

```go
package main

import (
	"github.com/surzia/go-stl/graph"
	"github.com/surzia/go-stl/hashmap"
	"github.com/surzia/go-stl/linkedlist"
	"github.com/surzia/go-stl/queue"
	"github.com/surzia/go-stl/set"
	"github.com/surzia/go-stl/stack"
	"github.com/surzia/go-stl/tree"
)

func main() {
	s := stack.NewStack()
	q := queue.NewQueue()
	e := set.NewSet()
	l := linkedlist.NewLinkedList()
	l := hashmap.NewHashMap()
	t := tree.NewTree()
	l := graph.NewGraph()
	...
}
```

## Todo
- [ ] Trie
- [ ] AVL Tree
- [ ] Circular Linked List
- [ ] Min-Max Heap
- [ ] LRU Cache
- [ ] Hash Mapped Array Trie