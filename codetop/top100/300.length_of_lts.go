package top100

func lengthOfLIS(nums []int) int {
	size := len(nums)
	dp := make([]int, size)
	ans := 1
	dp[0] = 1
	for i := 1; i < size; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}
