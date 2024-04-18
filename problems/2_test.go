package problems

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	// Test cases
	// Example 1: (2 -> 4 -> 3) + (5 -> 6 -> 4) = 7 -> 0 -> 8
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	want := &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}}
	testAddTwoNumbers(t, l1, l2, want)

	// Additional test cases can be added here.
}

func testAddTwoNumbers(t *testing.T, l1 *ListNode, l2 *ListNode, want *ListNode) {
	got := addTwoNumbers(l1, l2)
	if !isEqual(got, want) {
		t.Errorf("addTwoNumbers(%v, %v) = %v; want %v", l1, l2, got, want)
	}
}

func isEqual(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
