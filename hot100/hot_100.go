package hot100

// 1.两数之和
func twoSum(nums []int, target int) []int {
	value_index_map := make(map[int]int)
	for index, value := range nums {
		value_index_map[value] = index
	}
	for index, value := range nums {
		other_index, ok := value_index_map[target-value]
		if ok && index != other_index {
			return []int{index, other_index}
		}
	}
	return nil
}

// 2.两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	c1, c2, cr := l1, l2, head
	var n1, n2, nr int
	carry := 0
	for c1 != nil || c2 != nil || carry != 0 {
		if c1 != nil {
			n1 = c1.Val
			c1 = c1.Next
		} else {
			n1 = 0
		}
		if c2 != nil {
			n2 = c2.Val
			c2 = c2.Next
		} else {
			n2 = 0
		}
		nr = n1 + n2 + carry
		if nr > 9 {
			nr %= 10
			carry = 1
		} else {
			carry = 0
		}
		cr.Next = &ListNode{nr, nil}
		cr = cr.Next
	}
	return head.Next
}
