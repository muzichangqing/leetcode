package codetop

func removeNthFromEndV2(head *ListNode, n int) *ListNode {
	vh := &ListNode{Next: head}

	slow, fast := vh, vh
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return vh.Next
}
