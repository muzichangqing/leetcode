package bilibili

func isValid(s string) bool {
	size := len(s)
	st := []byte{' '}
	for i := 0; i < size; i++ {
		b := s[i]
		switch b {
		case '(', '[', '{':
			st = append(st, b)
		case ')':
			if st[len(st)-1] != '(' {
				return false
			}
			st = st[:len(st)-1]
		case ']':
			if st[len(st)-1] != '[' {
				return false
			}
			st = st[:len(st)-1]
		case '}':
			if st[len(st)-1] != '{' {
				return false
			}
			st = st[:len(st)-1]
		}
	}
	if len(st) == 1 {
		return true
	}
	return false
}
