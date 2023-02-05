package interviewtop

import "strconv"

func countAndSay(n int) string {
	b := []byte{'1'}
	for i := 1; i < n; i++ {
		nb := []byte{}
		j, k := 0, 1
		for j < len(b) {
			for ; k < len(b) && b[j] == b[k]; k++ {
			}
			nb = append(nb, strconv.Itoa(k-j)...)
			nb = append(nb, b[j])
			j, k = k, k+1
		}
		b = nb
	}
	return string(b)
}
