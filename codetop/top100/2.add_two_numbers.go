package top100

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	prev := dummyHead
	carry := 0
	for l1 != nil && l2 != nil {
		prev.Next = &ListNode{}
		sum := l1.Val + l2.Val + carry
		prev.Next.Val = sum % 10
		carry = sum / 10
		prev, l1, l2 = prev.Next, l1.Next, l2.Next
	}
	for l1 != nil {
		prev.Next = &ListNode{}
		sum := l1.Val + carry
		prev.Next.Val = sum % 10
		carry = sum / 10
		prev, l1 = prev.Next, l1.Next
	}
	for l2 != nil {
		prev.Next = &ListNode{}
		sum := l2.Val + carry
		prev.Next.Val = sum % 10
		carry = sum / 10
		prev, l2 = prev.Next, l2.Next
	}
	if carry > 0 {
		prev.Next = &ListNode{Val: carry}
	}
	return dummyHead.Next
}
