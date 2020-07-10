func solveSudoku(board [][]byte) {
	backtrack(board, 0, 0)
}
func backtrack(board [][]byte, row, col int) bool {
	n := 9
	if col == n { // 列满了，进入下一行
		return backtrack(board, row+1, 0)
	}
	if row == n { // 行满了，找到结果了
		return true
	}
	if board[row][col] != '.' { // 有预设值，不用填充
		return backtrack(board, row, col+1)
	}
	var c byte
	for c = '1'; c <= '9'; c++ {
		if !isValid(board, row, col, c) {
			continue
		}
		board[row][col] = c               // 做选择
		if backtrack(board, row, col+1) { // 填充下一列
			return true
		}
		board[row][col] = '.'
	}
	return false
}
func isValid(board [][]byte, row, col int, c byte) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == c || board[row][i] == c || board[row/3*3+i/3][col/3*3+i%3] == c {
			return false
		}
	}
	return true
}
