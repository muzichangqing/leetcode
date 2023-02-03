package bilibili

func longestValidParentheses(s string) int {
	dp := make([]int, len(s))
	maxLen := 0
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i-2 >= 0 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else {
				j := i - dp[i-1] - 1
				if j >= 0 && s[j] == '(' {
					if j-1 >= 0 {
						dp[i] = dp[j-1] + dp[i-1] + 2
					} else {
						dp[i] = dp[i-1] + 2
					}
				}
			}
			if dp[i] > maxLen {
				maxLen = dp[i]
			}
		}
	}
	return maxLen
}
