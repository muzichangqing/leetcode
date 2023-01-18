package interviewtop

func romanToInt(s string) int {
	prev := 0
	rs := 0
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for _, b := range []byte(s) {
		i := m[b]
		if i > prev {
			rs = rs + i - prev - prev
		} else {
			rs = rs + i
		}
		prev = i
	}
	return rs
}
