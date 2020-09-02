package leetcode

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	cursor1 := l1
	cursor2 := l2
	head := &ListNode{}
	current := head

	for cursor1 != nil && cursor2 != nil {
		if cursor1.Val > cursor2.Val {
			current.Next = cursor2
			cursor2 = cursor2.Next
		} else {
			current.Next = cursor1
			cursor1 = cursor1.Next
		}
		current = current.Next
	}
	if cursor1 != nil {
		current.Next = cursor1
	}
	if cursor2 != nil {
		current.Next = cursor2
	}

	return head.Next
}
