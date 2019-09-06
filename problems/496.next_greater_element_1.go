package leetcode

func nextGreaterElement(findNums []int, nums []int) []int {
	res := make([]int, len(findNums))
	tmpMap := make(map[int]int)
	numSize := len(nums)
	for fi := 0; fi < numSize; fi++ {
		tmpMap[nums[fi]] = -1
		for si := fi + 1; si < numSize; si++ {
			if nums[si] > nums[fi] {
				tmpMap[nums[fi]] = nums[si]
				break
			}
		}
	}
	for index, value := range findNums {
		res[index] = tmpMap[value]
	}
	return res
}
