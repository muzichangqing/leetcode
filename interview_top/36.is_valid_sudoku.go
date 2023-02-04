package interviewtop

func isValidSudoku(board [][]byte) bool {
	m := make(map[byte][][]int)
	for i, row := range board {
		for j, v := range row {
			if v == '.' {
				continue
			}
			pl := m[v]
			for _, p := range pl {
				if i == p[0] || j == p[1] {
					return false
				}
				if p[0]/3 == i/3 && p[1]/3 == j/3 {
					return false
				}
			}
			m[v] = append(m[v], []int{i, j})
		}
	}
	return true
}
