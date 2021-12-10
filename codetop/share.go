package codetop

// 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func createLinkedList(nums []int) *ListNode {
	head := &ListNode{}
	cursor := head
	for _, num := range nums {
		cursor.Next = &ListNode{Val: num}
		cursor = cursor.Next
	}
	return head.Next
}

func createSliceFromLinkedList(head *ListNode) []int {
	res := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		res = append(res, cur.Val)
	}
	return res
}

func listNodeEqual(l1 *ListNode, l2 *ListNode) bool {
	c1, c2 := l1, l2
	for c1 != nil && c2 != nil {
		if c1.Val != c2.Val {
			return false
		}
		c1 = c1.Next
		c2 = c2.Next
	}
	if c1 != nil || c2 != nil {
		return false
	}
	return true
}

func sliceEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
