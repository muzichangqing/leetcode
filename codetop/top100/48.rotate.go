package top100

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < (n+1)>>1; i++ {
		for j := i; j < n-1-i; j++ {
			i2 := n - 1 - i
			j2 := n - 1 - j
			matrix[i][j], matrix[j][i2], matrix[i2][j2], matrix[j2][i] =
				matrix[j2][i], matrix[i][j], matrix[j][i2], matrix[i2][j2]
		}
	}
}
