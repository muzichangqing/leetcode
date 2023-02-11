package top100

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	dummy := &ListNode{Next: head}
	prev := dummy
	i := 0
	for ; i < left-1; i++ {
		prev = prev.Next
	}
	tail := prev
	for ; i < right; i++ {
		tail = tail.Next
	}
	next := tail.Next
	tail.Next = nil
	reversedHead := reverseList(prev.Next)
	prev.Next.Next = next
	prev.Next = reversedHead
	return dummy.Next
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev, head = head, tmp
	}
	return prev
}
