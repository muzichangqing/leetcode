package top100

func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	if m == 0 || n == 0 {
		return 0
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	max2 := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	ans := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}
			ans = max2(ans, dp[i][j])
		}
	}
	return ans
}
