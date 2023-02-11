package top100

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			ptr := head
			for ptr != slow {
				ptr = ptr.Next
				slow = slow.Next
			}
			return ptr
		}
	}

	return nil
}
