package interviewtop

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	var isOdd bool
	if (m+n)%2 != 0 {
		isOdd = true
	}
	// 保证nums1是长度较小的数组
	if m > n {
		m, n, nums1, nums2 = n, m, nums2, nums1
	}
	if m == 0 {
		middle := n / 2
		if isOdd {
			return float64(nums2[middle])
		} else {
			return float64(nums2[middle-1]+nums2[middle]) / 2.0
		}
	}
	// 左边元素的个数
	totalLeft := (m + n + 1) / 2

	// 二分查找
	left, right := 0, m
	for left <= right {
		middle1 := (left + right) / 2
		middle2 := totalLeft - middle1

		// 判断是否符合条件
		if m == n {
			// 两个数组等长，可能出现middle1, middle2位于数组的边界
			if middle1 <= 0 {
				if nums1[0] >= nums2[n-1] {
					return float64((nums1[0] + nums2[n-1])) / 2.0
				} else {
					left = middle1 + 1
				}
			} else if middle2 <= 0 {
				if nums1[m-1] <= nums2[0] {
					return float64((nums1[m-1] + nums2[0])) / 2.0
				} else {
					right = middle1 - 1
				}
			}
		}
		if middle1 <= 0 {
			// 第一个数组左侧没有值
			if nums1[0] >= nums2[middle2-1] {
				if isOdd {
					return float64(nums2[middle2-1])
				} else {
					return float64(nums2[middle2-1]+min(nums1[0], nums2[middle2])) / 2.0
				}
			} else {
				left = middle1 + 1
			}
		} else if middle1 >= m {
			// 第一个数组右侧没有值
			if nums1[m-1] <= nums2[middle2] {
				if isOdd {
					return float64(max(nums1[m-1], nums2[middle2-1]))
				} else {
					return float64(nums2[middle2]+max(nums1[m-1], nums2[middle2-1])) / 2.0
				}
			} else {
				right = middle1 - 1
			}
		} else {
			// 一般情况
			if nums1[middle1-1] <= nums2[middle2] && nums1[middle1] > nums2[middle2-1] {
				if isOdd {
					return float64(max(nums1[middle1-1], nums2[middle2-1]))
				} else {
					return float64(max(nums1[middle1-1], nums2[middle2-1])+min(nums1[middle1], nums2[middle2])) / 2.0
				}
			} else {
				if nums1[middle1-1] > nums2[middle2] {
					right = middle1 - 1
				} else {
					left = middle1 + 1
				}
			}
		}

	}
	return 0
}
