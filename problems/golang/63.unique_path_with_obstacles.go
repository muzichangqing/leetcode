package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid == nil || len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	rowSize := len(obstacleGrid)
	colSize := len(obstacleGrid[0])
	dp := make([][]int, rowSize)
	for i := range dp {
		dp[i] = make([]int, colSize)
	}

	for i := 0; i < rowSize && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < colSize && obstacleGrid[0][j] == 0; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < rowSize; i++ {
		for j := 1; j < colSize; j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[rowSize-1][colSize-1]
}
