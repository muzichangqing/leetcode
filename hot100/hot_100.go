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
	return nil
}
