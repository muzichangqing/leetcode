package everyday

// 1375. 二进制字符串前缀一致的次数
func numTimesAllBlue(flips []int) int {
	n, ans, right := len(flips), 0, 0
	for i := 0; i < n; i++ {
		if flips[i] > right {
			right = flips[i]
		}
		if right == i+1 {
			ans++
		}
	}
	return ans
}
