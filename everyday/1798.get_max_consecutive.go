package everyday

import "sort"

func getMaximumConsecutive(coins []int) int {
	sort.Slice(coins, func(x, y int) bool {
		return coins[x] < coins[y]
	})
	res := 1
	for _, v := range coins {
		if v > res {
			break
		}
		res += v
	}
	return res
}
