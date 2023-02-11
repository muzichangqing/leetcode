package top100

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	head := lists[0]
	for i := 1; i < len(lists); i++ {
		head = mergeTwoLists(head, lists[i])
	}
	return head
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	prev := dummy
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			prev.Next = l2
			l2 = l2.Next
		} else {
			prev.Next = l1
			l1 = l1.Next
		}
		prev = prev.Next
	}
	if l1 != nil {
		prev.Next = l1
	} else {
		prev.Next = l2
	}
	return dummy.Next
}
