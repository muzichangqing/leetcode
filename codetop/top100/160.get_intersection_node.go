package top100

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB
	for pA != pB {
		pA = pA.Next
		pB = pB.Next
		if pA == nil && pB == nil {
			return nil
		}
		if pA == nil {
			pA = headB
		}
		if pB == nil {
			pB = headA
		}
	}
	return pA
}
