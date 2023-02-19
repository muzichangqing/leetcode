package top100

func searchRange(nums []int, target int) []int {
	n := len(nums)
	ans := []int{-1, -1}
	if n == 0 {
		return []int{-1, -1}
	}
	l, r := 0, n-1
	for l < r {
		mid := (l + r) >> 1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if r < 0 || nums[r] != target {
		return ans
	}
	ans[0] = r
	l, r = 0, n-1
	for l < r {
		mid := (l + r + 1) >> 1
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}
	ans[1] = l
	return ans
}
