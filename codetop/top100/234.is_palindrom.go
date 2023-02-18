package top100

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		slow = slow.Next
	}
	slow = reverse(slow)
	for slow != nil {
		if slow.Val != head.Val {
			return false
		}
		slow, head = slow.Next, head.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev, head = head, tmp
	}
	return prev
}
