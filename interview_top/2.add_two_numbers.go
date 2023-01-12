package interviewtop

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	flag := 0
	cursor := head
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + flag
		if val > 9 {
			val = val - 10
			flag = 1
		} else {
			flag = 0
		}
		cursor.Next = &ListNode{
			Val: val,
		}
		cursor, l1, l2 = cursor.Next, l1.Next, l2.Next
	}
	for l1 != nil {
		val := l1.Val + flag
		if val > 9 {
			val = val - 10
			flag = 1
		} else {
			flag = 0
		}
		cursor.Next = &ListNode{
			Val: val,
		}
		cursor, l1 = cursor.Next, l1.Next
	}
	for l2 != nil {
		val := l2.Val + flag
		if val > 9 {
			val = val - 10
			flag = 1
		} else {
			flag = 0
		}
		cursor.Next = &ListNode{
			Val: val,
		}
		cursor, l2 = cursor.Next, l2.Next
	}
	if flag == 1 {
		cursor.Next = &ListNode{
			Val: flag,
		}
	}
	return head.Next
}
