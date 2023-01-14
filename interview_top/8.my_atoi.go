package interviewtop

func myAtoi(s string) int {
	x, i, l := 0, 0, len(s)
	for ; i < l && s[i] == ' '; i++ {
	}
	if i == l {
		return 0
	}
	sign := 1
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}
	for ; i < l && s[i] == '0'; i++ {
	}
	for ; i < l && s[i] >= '0' && s[i] <= '9'; i++ {
		if x > 214748364 {
			if sign == 1 {
				return 2147483647
			} else {
				return -2147483648
			}
		}
		if x == 214748364 {
			if sign == 1 && s[i] > '7' {
				return 2147483647
			}
			if sign == -1 && s[i] > '8' {
				return -2147483648
			}
		}
		x = x*10 + int(s[i]-'0')
	}
	return sign * x
}
