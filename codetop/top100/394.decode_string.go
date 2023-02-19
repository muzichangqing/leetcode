package top100

func decodeString(s string) string {
	stack := []byte{}
	n := len(s)
	for i := 0; i < n; {
		switch {
		case s[i] >= '0' && s[i] <= '9':
			for ; s[i] >= '0' && s[i] <= '9'; i++ {
				stack = append(stack, s[i])
			}
		case s[i] == '[':
			stack = append(stack, s[i])
			i++
		case s[i] >= 'a' && s[i] <= 'z':
			for ; i < n && s[i] >= 'a' && s[i] <= 'z'; i++ {
				stack = append(stack, s[i])
			}
		case s[i] == ']':
			i++
			top := len(stack) - 1
			sub := []byte{}
			for ; stack[top] != '['; top-- {
				sub = append(sub, stack[top])
			}
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-1-i] = sub[len(sub)-1-i], sub[i]
			}
			top--
			count := 0
			multi := 1
			for ; top >= 0 && stack[top] >= '0' && stack[top] <= '9'; top-- {
				count = int(stack[top]-'0')*multi + count
				multi *= 10
			}
			stack = stack[:top+1]
			for ; count > 0; count-- {
				stack = append(stack, sub...)
			}
		}
	}
	return string(stack)
}
