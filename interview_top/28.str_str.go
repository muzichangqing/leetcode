package interviewtop

func strStr(haystack string, needle string) int {
	nh, nn := len(haystack), len(needle)
	switch {
	case nn == 0:
		return 0
	case nn == 1:
		return findByte(haystack, needle[0])
	case nn == nh:
		if haystack == needle {
			return 0
		}
		return -1
	case nn > nh:
		return -1
	}
	t := nh - nn + 1
	b0 := needle[0]
	i := findByte(haystack, b0)
	if i == -1 {
		return i
	}
	for i < t {
		if haystack[i:i+nn] == needle {
			return i
		}
		j := findByte(haystack[i+1:], b0)
		if j == -1 {
			return j
		}
		i += j + 1
	}
	return -1
}

func findByte(haystack string, b byte) int {
	for i := 0; i < len(haystack); i++ {
		if b == haystack[i] {
			return i
		}
	}
	return -1
}
