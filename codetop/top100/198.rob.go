package top100

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	dp0, dp1 := nums[0], max(nums[0], nums[1])
	dp := dp1
	for i := 2; i < n; i++ {
		dp = max(dp0+nums[i], dp1)
		dp0, dp1 = dp1, dp
	}
	return dp
}
