package top100

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x, y, z int) int {
		rs := x
		if y < rs {
			rs = y
		}
		if z < rs {
			rs = z
		}
		return rs
	}
	rs := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			if j == 0 || i == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
			rs = max(dp[i][j], rs)
		}
	}
	return rs * rs
}
