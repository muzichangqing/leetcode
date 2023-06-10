package everyday

import "sort"

// 1170. 比较字符串最小字母出现频次
func numSmallerByFrequency(queries []string, words []string) []int {
	foo := func(s string) int {
		size := len(s)
		if size == 0 {
			return 0
		}
		count := 1
		smaller := s[0]
		for i := 1; i < size; i++ {
			c := s[i]
			if c > smaller {
				continue
			}
			if c == smaller {
				count++
				continue
			}
			count, smaller = 1, c
		}
		return count
	}
	wfs := make([]int, len(words))
	for i, w := range words {
		wfs[i] = foo(w)
	}
	sort.Slice(wfs, func(i, j int) bool {
		return wfs[i] > wfs[j]
	})
	nums := make([]int, len(queries))
	for i, w := range queries {
		num := 0
		f := foo(w)
		for _, wf := range wfs {
			if wf > f {
				num++
			} else {
				break
			}
		}
		nums[i] = num
	}
	return nums
}
