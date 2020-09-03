package leetcode

var result [][]string

func solveNQueens(n int) [][]string {
	result = [][]string{}
	queens := make([]int, n)
	for k := range queens {
		queens[k] = -1
	}
	columns := make(map[int]bool)
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}
	backtrace(queens, n, 0, columns, diagonals1, diagonals2)
	return result
}

func backtrace(queens []int, n, row int, columns, diagonals1, diagnoals2 map[int]bool) {
	if row == n {
		result = append(result, generateBoard(queens, n))
		return
	}
	for i := 0; i < n; i++ {
		if columns[i] { // i 列上存在queen
			continue
		}
		diagnoal1 := row - i // 左上 到 右下的斜线，如果是同一条斜线的坐标，则 row - col值相等
		if diagonals1[diagnoal1] {
			continue
		}
		diagnoal2 := row + i // 右上 到 左下的斜线， 如果是同一条斜线， 则 row + col 值相等
		if diagnoals2[diagnoal2] {
			continue
		}
		queens[row] = i // 第 row 行 i 列 上添加一个queen
		columns[i] = true
		diagonals1[diagnoal1] = true
		diagnoals2[diagnoal2] = true
		backtrace(queens, n, row+1, columns, diagonals1, diagnoals2) // 下一行

		// 回溯到未添加queue的状态
		queens[row] = -1
		columns[i] = false
		diagonals1[diagnoal1] = false
		diagnoals2[diagnoal2] = false
	}
}

func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}
