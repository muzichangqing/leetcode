package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	firstHalfEnd := endOfFirstHalf(head)
	secondHalfHead := reverseLinkedList(firstHalfEnd.Next)

	p1, p2 := head, secondHalfHead
	result := true
	for p2 != nil {
		if p1.Val != p2.Val {
			result = false
			break
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	firstHalfEnd.Next = reverseLinkedList(secondHalfHead)

	return result
}

func reverseLinkedList(head *ListNode) *ListNode {
	var cursor, pre *ListNode
	cursor = head

	for cursor != nil {
		tmp := cursor.Next
		cursor.Next = pre
		pre = cursor
		cursor = tmp
	}

	return pre
}

func endOfFirstHalf(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
