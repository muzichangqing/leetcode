package leetcode

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	fakerHead := &ListNode{Next: list1}
	cursor := fakerHead
	for i := 0; i < a; i++ {
		cursor = cursor.Next
	}
	prev := cursor
	for i := 0; i < b-a+1; i++ {
		cursor = cursor.Next
	}
	tail := cursor.Next
	prev.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = tail
	return fakerHead.Next
}
