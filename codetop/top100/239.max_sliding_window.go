package top100

import (
	"container/heap"
	"sort"
)

var a []int

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool {
	return a[h.IntSlice[i]] > a[h.IntSlice[j]]
}

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}

func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	n := len(nums)
	h := &hp{
		IntSlice: make([]int, k),
	}
	for i := 0; i < k; i++ {
		h.IntSlice[i] = i
	}
	heap.Init(h)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[h.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(h, i)
		for h.IntSlice[0] <= i-k {
			heap.Pop(h)
		}
		ans = append(ans, nums[h.IntSlice[0]])
	}
	return ans
}
