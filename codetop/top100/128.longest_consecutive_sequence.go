package top100

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}

	ans := 0
	for k := range m {
		if m[k-1] {
			continue
		}
		consecutive := 0
		for i := k; m[i]; i++ {
			consecutive++
		}
		if ans < consecutive {
			ans = consecutive
		}
	}
	return ans
}
