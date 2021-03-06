# Golang 数据结构和算法

这个仓库包含许多常见的数据结构和算法的 Golang 的实现示例。

## 数据结构

数据结构是一种在计算机中组织和存储数据的特殊方式，以便可以有效地访问和修改数据。更准确地说，数据结构是数据值、它们之间的关系以及可应用于数据的功能或操作的集合。

已经实现的数据结构：

- [集合](collections/set/set.go)
- [栈](collections/stack/stack.go)
- [队列](collections/queue/queue.go)
- [哈希表](collections/hashmap/hashmap.go)
- [链表](collections/linkedlist/linkedlist.go)
- [图](collections/graph/graph.go)
- [二叉搜索树](collections/tree/bin_search_tree.go)
- [前缀树](collections/tree/trie.go)
- [环形缓存区](collections/ringbuffer/ring_buffer.go)
- [LRU Cache](collections/cache//lru_cache.go)

## 算法

算法是如何解决一类问题的明确规范，它是一组精确定义操作序列的规则。

已经实现的算法：

### 查找
- [二分查找](algorithms/search/binary_search.go)
- [插值查找](algorithms/search/interpolation_search.go)
### 排序
- [冒泡排序](algorithms/sort/bubble_sort.go)
- [插入排序](algorithms/sort/insertion_sort.go)
- [归并排序](algorithms/sort/merge_sort.go)
- [快速排序](algorithms/sort/quick_sort.go)
- [选择排序](algorithms/sort/selection_sort.go)
- [希尔排序](algorithms/sort/shell_sort.go)

### 数学相关
- [斐波那契数列](algorithms/mathematical/fibonacci.go)

## 应用

这里记录了一些有趣的Go实现的应用：
- [go实现lolcat命令](applications/gololcat/main.go)
