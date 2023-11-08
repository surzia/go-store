package collections

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}

// 向二叉树中插入节点
func (b *BinaryTree) Insert(value int) {
	if b.Root == nil {
		b.Root = &TreeNode{Value: value, Left: nil, Right: nil}
	} else {
		b.Root.insertTreeNode(value)
	}
}

// 递归地在树中找到合适的位置插入新节点
func (n *TreeNode) insertTreeNode(value int) {
	if value < n.Value {
		if n.Left == nil {
			n.Left = &TreeNode{Value: value, Left: nil, Right: nil}
		} else {
			n.Left.insertTreeNode(value)
		}
	} else if value > n.Value {
		if n.Right == nil {
			n.Right = &TreeNode{Value: value, Left: nil, Right: nil}
		} else {
			n.Right.insertTreeNode(value)
		}
	}
}
