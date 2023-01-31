package interviewtop

func mergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode
	for _, list := range lists {
		head = mergeTwoLists(head, list)
	}
	return head
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	head := &ListNode{}
	cursor := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cursor.Next = list1
			list1 = list1.Next
		} else {
			cursor.Next = list2
			list2 = list2.Next
		}
		cursor = cursor.Next
	}
	if list1 != nil {
		cursor.Next = list1
	} else if list2 != nil {
		cursor.Next = list2
	}
	return head.Next
}
