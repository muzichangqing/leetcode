package interviewtop

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	m := 0
	for l <= r {
		m = (l + r) >> 1
		if nums[m] == target {
			return m
		}
		if nums[m] >= nums[l] {
			if target >= nums[l] && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if target >= nums[l] || target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
	}
	return -1
}

/*
4 5 6 7 8 9 0 1 2 3
二分查找 分割后的分析
1. 4 5 6
	普通情况，直接判断在不在这个区间
2. 7 8 9 0
	如果 > 7 则在这个区间查找
	如果 < 0 在这个区间查找
不在这连个区间 返回 -1
*/
