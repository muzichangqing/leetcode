package leetcode

func islandPerimeter(grid [][]int) int {
	perimeter := 0
	rows := len(grid)
	cols := len(grid[0])
	for rowIndex, rowValue := range grid {
		for colIndex, colValue := range rowValue {
			if colValue == 0 {
				continue
			}
			perimeter += 4
			if colIndex < cols-1 && rowValue[colIndex+1] == 1 {
				perimeter -= 2
			}
			if rowIndex < rows-1 && grid[rowIndex+1][colIndex] == 1 {
				perimeter -= 2
			}
		}
	}
	return perimeter
}
