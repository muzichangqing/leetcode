package top100

func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	dp_0, dp_1 := 0, -prices[0]

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := 1; i < n; i++ {
		dp_0 = max(dp_0, dp_1+prices[i])
		dp_1 = max(dp_1, dp_0-prices[i])
	}
	return dp_0
}
