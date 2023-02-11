package top100

func permute(nums []int) [][]int {
	numsSize := len(nums)
	ansSize := 1
	for i := 1; i <= numsSize; i++ {
		ansSize *= i
	}
	ans := make([][]int, 0, ansSize)
	var backtrace func(int)
	backtrace = func(first int) {
		if first == numsSize {
			newNums := make([]int, numsSize)
			copy(newNums, nums)
			ans = append(ans, newNums)
			return
		}
		for i := first; i < numsSize; i++ {
			nums[i], nums[first] = nums[first], nums[i]
			backtrace(first + 1)
			nums[i], nums[first] = nums[first], nums[i]
		}
	}
	backtrace(0)
	return ans
}
