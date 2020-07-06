package leetcode

func searchInsert(nums []int, target int) int {
	for k, v := range nums {
		if v == target {
			return k
		}
		if v > target {
			return k
		}
	}
	return len(nums)
}
