package interviewtop

func lengthOfLongestSubstring(s string) int {
	index := make(map[byte]bool)
	l, r, max := 0, 0, 0
	for ; r < len(s); r++ {
		if index[s[r]] {
			for ; s[r] != s[l]; l++ {
				index[s[l]] = false
			}
			l++
		} else {
			index[s[r]] = true
			if r-l+1 > max {
				max = r - l + 1
			}
		}
	}
	return max
}
