package codetop

func isValid(s string) bool {
	size := len(s)
	stack := make([]byte, size)
	top := -1
	for i := 0; i < size; i++ {
		b := s[i]
		switch b {
		case '(', '[', '{':
			top++
			stack[top] = b
		case ')', ']', '}':
			if top == -1 {
				return false
			}
			l := stack[top]
			if l == '(' && b == ')' || l == '[' && b == ']' || l == '{' && b == '}' {
				top--
			} else {
				return false
			}
		}
	}
	return top == -1
}
