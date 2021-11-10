package codetop

import "math"

//8. 字符串转换整数(atoi)
func myAtoi(s string) int {
	const (
		START = iota
		SIGN
		IN_NUMBER
		END
	)
	states := [][]int{
		{START, SIGN, IN_NUMBER, END},
		{END, END, IN_NUMBER, END},
		{END, END, IN_NUMBER, END},
		{END, END, END, END},
	}

	state := START
	sign := 1
	ans := 0
	getCol := func(c rune) int {
		if c == ' ' {
			return 0
		}
		if c == '+' || c == '-' {
			return 1
		}
		if c >= '0' && c <= '9' {
			return 2
		}
		return 3
	}
	for _, c := range s {
		state = states[state][getCol(c)]
		if state == END {
			break
		}
		if state == SIGN && c == '-' {
			sign = -1
		} else if state == IN_NUMBER {
			ans = ans*10 + int(c-'0')
			if ans > math.MaxInt32 {
				if sign == 1 {
					return math.MaxInt32
				} else {
					return math.MinInt32
				}
			}
		}
	}
	return ans * sign
}

// 2. 两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	c1, c2, c3 := l1, l2, head
	carry := 0
	for c1 != nil && c2 != nil {
		val := c1.Val + c2.Val + carry
		c3.Next = &ListNode{Val: val % 10}
		carry = val / 10
		c1 = c1.Next
		c2 = c2.Next
		c3 = c3.Next
	}
	for c1 != nil {
		val := c1.Val + carry
		c3.Next = &ListNode{Val: val % 10}
		carry = val / 10
		c1 = c1.Next
		c3 = c3.Next
	}
	for c2 != nil {
		val := c2.Val + carry
		c3.Next = &ListNode{Val: val % 10}
		carry = val / 10
		c2 = c2.Next
		c3 = c3.Next
	}
	if carry == 1 {
		c3.Next = &ListNode{Val: 1}
	}
	return head.Next
}

// 82. 删除排序链表中的重复元素 II
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev, corrent := dummy, head
	for corrent != nil {
		next := corrent.Next
		for next != nil && next.Val == corrent.Val {
			next = next.Next
		}
		if next != corrent.Next {
			prev.Next = next
			corrent = next
		} else {
			prev = corrent
			corrent = corrent.Next
		}
	}
	return dummy.Next
}

// 144. 二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {
	ans := []int{}

	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return ans
}

// 148.排序链表
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	merge := func(h1, h2 *ListNode) *ListNode {
		h3 := &ListNode{}
		c1, c2, c3 := h1, h2, h3
		for c1 != nil && c2 != nil {
			if c1.Val > c2.Val {
				c3.Next = c2
				c2 = c2.Next
			} else {
				c3.Next = c1
				c1 = c1.Next
			}
			c3 = c3.Next
		}
		if c1 != nil {
			c3.Next = c1
		} else if c2 != nil {
			c3.Next = c2
		}

		return h3.Next
	}

	var mergeSort func(*ListNode) *ListNode
	mergeSort = func(head *ListNode) *ListNode {
		if head.Next == nil {
			return head
		}
		slow, fast := head, head.Next
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}
		h1 := mergeSort(slow.Next)
		slow.Next = nil
		h2 := mergeSort(head)
		return merge(h1, h2)
	}
	return mergeSort(head)
}
