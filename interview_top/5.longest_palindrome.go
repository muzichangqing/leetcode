package interviewtop

func longestPalindrome(s string) string {
	sl := len(s)
	if sl == 0 {
		return s
	}
	expand := func(l, r int) int {
		for l >= 0 && r < sl && s[l] == s[r] {
			l = l - 1
			r = r + 1
		}
		return r - l - 1
	}
	start, end := 0, 0
	for i := 0; i < sl; i++ {
		len1, len2 := expand(i, i), expand(i, i+1)
		mLen := max(len1, len2)
		if mLen > end-start+1 {
			start = i - (mLen-1)/2
			end = i + mLen/2
		}
	}
	return string(s[start : end+1])
}
