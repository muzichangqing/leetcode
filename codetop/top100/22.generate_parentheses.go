package top100

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	parenthes := make([]byte, 0)
	var backtrace func(left, right, size int)
	backtrace = func(left, right, size int) {
		if left == 0 && right == 0 {
			ans = append(ans, string(parenthes))
			return
		}
		if left < 0 || left > right {
			return
		}
		parenthes = append(parenthes, '(')
		backtrace(left-1, right, size+1)
		parenthes = parenthes[:size]
		if left < right {
			parenthes = append(parenthes, ')')
			backtrace(left, right-1, size+1)
			parenthes = parenthes[:size]
		}
	}
	backtrace(n, n, 0)
	return ans
}
