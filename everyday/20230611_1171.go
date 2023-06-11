package everyday

// 1171. 从链表中删去总和值为零的连续节点
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
	fh := &ListNode{Next: head}
	pre := fh

	for pre.Next != nil {
		cur := pre.Next
		sum := 0
		for cur != nil {
			sum += cur.Val
			if sum == 0 {
				pre.Next = cur.Next
				cur = pre.Next
			} else {
				cur = cur.Next
			}
		}
		pre = pre.Next
	}

	return fh.Next
}
