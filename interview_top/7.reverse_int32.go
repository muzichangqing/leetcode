package interviewtop

func reverse(x int) int {
	var y int
	for i := 0; i < 9 && x != 0; i++ {
		y = y*10 + x%10
		x /= 10
	}
	if x == 0 {
		return y
	}
	if y > 214748364 || y < -214748364 {
		return 0
	}
	return y*10 + x
}
