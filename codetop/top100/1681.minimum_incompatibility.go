package top100

import (
	"math"
	"sort"
)

func minimumIncompatibility(nums []int, k int) int {
	size := len(nums)
	per := size / k
	if per == 0 || per == 1 {
		return 0
	}
	counter := make(map[int]int)
	keys := []int{}
	for _, v := range nums {
		if counter[v] == 0 {
			keys = append(keys, v)
		}
		counter[v] += 1
		if counter[v] > k {
			return -1
		}
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	ans := math.MaxInt32
	var backtrace func(idx, curNum, delta int, curList []int)
	backtrace = func(idx, curNum, delta int, curList []int) {
		if delta >= ans {
			return
		}
		if len(curList) == per {
			delta += curList[len(curList)-1] - curList[0]
			curList = curList[:0]
			curNum = -1
		}
		if idx == size {
			if ans < delta {
				ans = delta
			}
			return
		}
		for _, i := range keys {
			if i <= curNum || counter[i] == 0 {
				continue
			}
			counter[i] -= 1
			curList = append(curList, i)
			backtrace(idx+1, i, delta, curList)
			curList = curList[:len(curList)-1]
			counter[i] += 1
			if len(curList) == 0 {
				break
			}
		}
	}
	backtrace(0, -1, 0, []int{})
	return ans
}
