package top100

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}
	k := (m + n + 1) >> 1
	left, right := 0, m
	for left < right {
		i := (left + right + 1) >> 1
		j := k - i
		if nums1[i-1] > nums2[j] {
			right = i - 1
		} else {
			left = i
		}
	}
	i, j := left, k-left
	nums1Left, nums2Left := math.MinInt32, math.MinInt32
	if i > 0 {
		nums1Left = nums1[i-1]
	}
	if j > 0 {
		nums2Left = nums2[j-1]
	}
	if (m+n)%2 == 1 {
		return float64(max(nums1Left, nums2Left))
	}
	nums1Right, nums2Right := math.MaxInt32, math.MaxInt32
	if i < m {
		nums1Right = nums1[i]
	}
	if j < n {
		nums2Right = nums2[j]
	}
	return float64(max(nums1Left, nums2Left)+min(nums1Right, nums2Right)) / 2.0
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
