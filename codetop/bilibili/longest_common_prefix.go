package bilibili

func longestCommonPrefix(strs []string) string {
	str1 := strs[0]
	n := len(str1)
	for k := 1; k < len(strs); k++ {
		i := 0
		for ; i < n && i < len(strs[k]); i++ {
			if str1[i] != strs[k][i] {
				break
			}
		}
		if i == 0 {
			return ""
		}
		n = i
	}
	return str1[:n]
}
