package leetcode

func checkXMatrix(grid [][]int) bool {
	n := len(grid) - 1
	for i, row := range grid {
		for j, v := range row {
			if i == j || i+j == n {
				if v == 0 {
					return false
				}
			} else {
				if v != 0 {
					return false
				}
			}
		}
	}
	return true
}
