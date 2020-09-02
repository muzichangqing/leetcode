package leetcode

func removeElement(nums []int, val int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}
	i := 0
	for ; i < size; i++ {
		if nums[i] == val {
			break
		}
	}
	for j := i + 1; j < size; j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}
