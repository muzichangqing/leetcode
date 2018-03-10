func islandPerimeter(grid [][]int) int {
    perimeter := 0
    rows := len(grid)
    cols := len(grid[0])
    for row_index, row_value := range(grid) {
        for col_index, col_value := range(row_value) {
            if col_value == 0 {
                continue
            }
            perimeter += 4
            if col_index < cols - 1 && row_value[col_index + 1] == 1 {
                perimeter -= 2
            }
            if row_index < rows - 1 && grid[row_index + 1][col_index] == 1 {
                perimeter -= 2
            }
        }
    }
    return perimeter
}
