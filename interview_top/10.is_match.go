package interviewtop

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	var isMatchDp func(int, int) bool
	isMatchDp = func(i, j int) bool {
		if dp[i][j] != 0 {
			return dp[i][j] == 1
		}
		match := false
		if j == n {
			match = i == m
		} else {
			firstMatch := i < m && (s[i] == p[j] || p[j] == '.')
			if j < n-1 && p[j+1] == '*' {
				match = isMatchDp(i, j+2) || firstMatch && isMatchDp(i+1, j)
			} else {
				match = firstMatch && isMatchDp(i+1, j+1)
			}
		}
		if match {
			dp[i][j] = 1
		} else {
			dp[i][j] = -1
		}
		return match
	}
	return isMatchDp(0, 0)
}
