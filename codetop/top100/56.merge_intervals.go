package top100

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := make([][]int, 0)
	it := intervals[0]
	for i := 1; i < len(intervals); i++ {
		it2 := intervals[i]
		if it2[0] > it[1] {
			ans = append(ans, it)
			it = it2
			continue
		}
		if it2[1] > it[1] {
			it[1] = it2[1]
		}
	}
	ans = append(ans, it)
	return ans
}
