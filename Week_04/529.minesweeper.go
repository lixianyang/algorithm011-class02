// 递归
/*
func updateBoard(board [][]byte, click []int) [][]byte {
    update(board, len(board), len(board[0]), click[0], click[1])
    return board
}
func update(board [][]byte, rows, cols, row, col int) {
    if row < 0 || col < 0 || row >= rows || col >= cols {
        return
    }
    if board[row][col] == 'M' {
        board[row][col] = 'X'
        return
    }
    if board[row][col] == 'E' {
        n := countM(board, rows, cols, row, col)
        if n == '0' {
            board[row][col] = 'B'
            for i := -1; i <= 1; i++ {
                for j := -1; j <= 1; j++ {
                    if i == 0 && j == 0 {
                        continue
                    }
                    update(board, rows, cols, row+i, col+j)
                }
            }
        } else {
            board[row][col] = n
        }
    }
    return
}
func countM(board [][]byte, rows, cols, row, col int) byte {
    var n byte = '0'
    for i := -1; i <= 1; i++ {
        for j := -1; j <= 1; j++ {
            currentRow, currentCol := row+i, col+j
            if currentRow < 0 || currentRow >= rows || currentCol < 0 || currentCol >= cols {
                continue
            }
            if board[currentRow][currentCol] == 'M' {
                n++
            }
        }
    }
    return n
}
*/
func updateBoard(board [][]byte, click []int) [][]byte {
	rows, cols := len(board), len(board[0])
	queue := [][]int{click}
	for len(queue) > 0 {
		row, col := queue[0][0], queue[0][1]
		queue = queue[1:]
		if board[row][col] == 'E' {
			board[row][col] = 'B'
			n := countM(board, rows, cols, row, col)
			if n == '0' {
				for i := -1; i <= 1; i++ {
					for j := -1; j <= 1; j++ {
						currentRow, currentCol := row+i, col+j
						if currentRow < 0 || currentRow >= rows || currentCol < 0 || currentCol >= cols || i == 0 && j == 0 {
							continue
						}
						if board[currentRow][currentCol] == 'E' {
							queue = append(queue, []int{currentRow, currentCol})
						}
					}
				}
			} else {
				board[row][col] = n
			}
		} else if board[row][col] == 'M' {
			board[row][col] = 'X'
			return board
		}
	}
	return board
}
func countM(board [][]byte, rows, cols, row, col int) byte {
	var n byte = '0'
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			currentRow, currentCol := row+i, col+j
			if currentRow < 0 || currentRow >= rows || currentCol < 0 || currentCol >= cols {
				continue
			}
			if board[currentRow][currentCol] == 'M' {
				n++
			}
		}
	}
	return n
}
