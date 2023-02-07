package everyday

import (
	"sort"
	"strconv"
	"strings"
)

func alertNames(keyName []string, keyTime []string) []string {
	m := make(map[string][]int)
	for i, v := range keyName {
		m[v] = append(m[v], str2Minutes(keyTime[i]))
	}
	for _, v := range m {
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
	}
	ans := make([]string, 0)

	for k, v := range m {
		i, j := 0, 0
		for ; j < len(v); j++ {
			if v[j] < v[i] {
				i = j
				continue
			}
			for ; v[j]-v[i] > 60; i++ {
			}
			if j-i >= 2 {
				ans = append(ans, k)
				break
			}
		}
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i] < ans[j]
	})
	return ans
}

func str2Minutes(s string) int {
	ss := strings.Split(s, ":")
	h, _ := strconv.Atoi(ss[0])
	m, _ := strconv.Atoi(ss[1])
	return h*60 + m
}
