package top100

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")
	if len(s) == 0 {
		return 0
	}
	sign := 1
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		s = s[1:]
		sign = -1
	}
	num := 0
	maxSub0 := math.MaxInt32 / 10
	maxSub1P := byte(math.MaxInt32 % 10)
	maxSub1N := byte(-1 * math.MinInt32 % 10)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			break
		}
		c -= '0'
		if num > maxSub0 {
			num = math.MaxInt32
			if sign == -1 {
				num = math.MinInt32
			}
			return num
		}
		if num == maxSub0 {
			if sign > 0 && c >= maxSub1P {
				return math.MaxInt32
			}
			if sign < 0 && c >= maxSub1N {
				return math.MinInt32
			}
		}
		num = num*10 + int(c)
	}
	return num * sign
}
