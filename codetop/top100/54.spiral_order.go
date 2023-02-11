package top100

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	total := m * n
	ans := make([]int, total)
	x, y := 0, 0
	d := 0
	for i := 0; i < total; i++ {
		ans[i] = matrix[x][y]
		visited[x][y] = true
		nx, ny := x+directions[d][0], y+directions[d][1]
		if nx < 0 || ny < 0 || nx >= m || ny >= n || visited[nx][ny] {
			d = (d + 1) % 4
			nx, ny = x+directions[d][0], y+directions[d][1]
		}
		x, y = nx, ny
	}
	return ans
}
