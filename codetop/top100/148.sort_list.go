package top100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	return mergeSortList(head, nil)
}

func mergeSortList(head, tail *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow, fast := head, head
	for fast != tail {
		fast = fast.Next
		slow = slow.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	return mergeTwoList(mergeSortList(head, slow), mergeSortList(slow, tail))
}

func mergeTwoList(h1, h2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	prev := dummyHead
	for h1 != nil && h2 != nil {
		if h1.Val < h2.Val {
			prev.Next = h1
			h1 = h1.Next
		} else {
			prev.Next = h2
			h2 = h2.Next
		}
		prev = prev.Next
	}
	if h1 != nil {
		prev.Next = h1
	} else {
		prev.Next = h2
	}
	return dummyHead.Next
}
