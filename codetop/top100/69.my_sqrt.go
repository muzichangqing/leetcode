package top100

func mySqrt(x int) int {
	l, r, ans := 0, x, 0
	for l <= r {
		mid := (l + r) >> 1
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
