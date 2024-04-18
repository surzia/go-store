package problems

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var front = &ListNode{
		Val:  -1,
		Next: head,
	}
	var back = head
	for n > 0 {
		back = back.Next
		n--
	}

	for back != nil {
		front = front.Next
		back = back.Next
	}

	if head == front.Next {
		return head.Next
	}
	front.Next = front.Next.Next
	return head
}
