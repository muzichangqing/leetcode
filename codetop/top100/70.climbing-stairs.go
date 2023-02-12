package top100

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	x, y := 1, 2
	for z := 3; z <= n; z++ {
		x, y = y, x+y
	}
	return y
}
