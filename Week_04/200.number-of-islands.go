func numIslands(grid [][]byte) int {
    if len(grid) == 0 { return 0 }
    rows, cols := len(grid), len(grid[0])
    count := 0
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if grid[i][j] == '1' {
                count++
                mask(grid, rows, cols, i, j)
            }
        }
    }
    return count
}
func mask(grid [][]byte, rows, cols, i, j int) {
    if i >= rows || j >= cols || i < 0 || j < 0 || grid[i][j] == '0' {
        return
    }
    grid[i][j] = '0'
    mask(grid, rows, cols, i-1, j)
    mask(grid, rows, cols, i+1, j)
    mask(grid, rows, cols, i, j-1)
    mask(grid, rows, cols, i, j+1)
}
