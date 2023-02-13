package top100

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Next: head}
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val != cur.Next.Next.Val {
			cur = cur.Next
			continue
		}
		x := cur.Next.Val
		for cur.Next != nil && cur.Next.Val == x {
			cur.Next = cur.Next.Next
		}
	}

	return dummyHead.Next
}
