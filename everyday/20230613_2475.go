package everyday

// 2475.数组中不等三元组的数目
func unequalTriplets(nums []int) int {
	size := len(nums)
	count := 0
	for i := 0; i < size-2; i++ {
		for j := i + 1; j < size-1; j++ {
			if nums[i] == nums[j] {
				continue
			}
			for k := j + 1; k < size; k++ {
				if nums[i] == nums[k] || nums[j] == nums[k] {
					continue
				}
				count++
			}
		}
	}
	return count
}
