package top100

func reorderList(head *ListNode) {
	var list []*ListNode
	for head != nil {
		list = append(list, head)
		head = head.Next
	}
	i, j := 0, len(list)-1
	for i < j {
		list[i].Next = list[j]
		i++
		if i == j {
			break
		}
		list[j].Next = list[i]
		j--
	}
	list[i].Next = nil
}
