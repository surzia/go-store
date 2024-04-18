package problems

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	node := head
	carry := 0
	for l1 != nil && l2 != nil {
		res := l1.Val + l2.Val + carry
		carry = res / 10
		res = res % 10
		node.Val = res
		l1 = l1.Next
		l2 = l2.Next
		if l1 != nil || l2 != nil {
			node.Next = &ListNode{}
			node = node.Next
		}
	}

	for l1 != nil {
		res := l1.Val + carry
		carry = res / 10
		res = res % 10
		node.Val = res
		l1 = l1.Next
		if l1 != nil {
			node.Next = &ListNode{}
			node = node.Next
		}
	}

	for l2 != nil {
		res := l2.Val + carry
		carry = res / 10
		res = res % 10
		node.Val = res
		l2 = l2.Next
		if l2 != nil {
			node.Next = &ListNode{}
			node = node.Next
		}
	}

	if carry != 0 {
		node.Next = &ListNode{Val: carry}
	}

	return head
}
