package top100

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	n := len(candidates)

	var backtrace func(int, int, []int, int)
	backtrace = func(idx int, itemsIdx int, items []int, res int) {
		if res == 0 {
			rs := make([]int, len(items))
			copy(rs, items)
			ans = append(ans, rs)
			return
		}
		if idx >= n {
			return
		}
		m := 0
		num := candidates[idx]
		items = items[:itemsIdx]
		for res >= 0 {
			backtrace(idx+1, itemsIdx+m, items, res)
			items = append(items, num)
			res -= num
			m += 1
		}
	}

	backtrace(0, 0, []int{}, target)

	return ans
}
