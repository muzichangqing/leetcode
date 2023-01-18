package interviewtop

func longestCommonPrefix(strs []string) string {
	r := len(strs[0])
	for i := 1; i < len(strs) && r > 0; i++ {
		if len(strs[i]) < r {
			r = len(strs[i])
		}
		for j := 0; j < r; j++ {
			if strs[i][j] != strs[0][j] {
				r = j
				break
			}
		}
	}
	return strs[0][:r]
}
