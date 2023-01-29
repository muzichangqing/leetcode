package interviewtop

import "sort"

func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	numsSize := len(nums)
	for i := 0; i < numsSize; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		k := numsSize - 1
		j := i + 1
		for j < k {
			if j != i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return result
}
