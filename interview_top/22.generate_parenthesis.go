package interviewtop

func generateParenthesis(n int) []string {
	var result []string

	var addP func(ln int, rn int, r []byte, ri int)
	addP = func(ln int, rn int, r []byte, ri int) {
		if ln == 0 && rn == 0 {
			result = append(result, string(r))
		} else {
			ri++
			if ln > 0 {
				r[ri] = '('
				addP(ln-1, rn, r, ri)
			}
			if rn > 0 && rn > ln {
				r[ri] = ')'
				addP(ln, rn-1, r, ri)
			}
		}
	}
	r := make([]byte, n*2)
	addP(n, n, r, -1)
	return result
}
