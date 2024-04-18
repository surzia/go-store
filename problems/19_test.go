package problems

import "testing"

func sliceToListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}

func listNodeToSlice(head *ListNode) []int {
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return nums
}

func TestRemoveNthFromEnd(t *testing.T) {
	testCases := []struct {
		head     []int
		n        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1}, 1, []int{}},
		{[]int{1, 2}, 1, []int{1}},
	}

	for _, tc := range testCases {
		head := sliceToListNode(tc.head)
		expected := tc.expected

		result := removeNthFromEnd(head, tc.n)
		resultSlice := listNodeToSlice(result)

		if !compareSlice(resultSlice, expected) {
			t.Errorf("removeNthFromEnd(%v, %d) = %v; want %v", tc.head, tc.n, resultSlice, expected)
		}
	}
}

func compareSlice(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
