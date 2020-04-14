package _7_sudoku_solver

//var rows, cols, cells [9]map[byte]bool

func solveSudoku(board [][]byte) {
	_ = dfs(board, 0)
}

func dfs(board [][]byte, index int) bool {
	if index > 80 {
		return true
	}
	row, col := index/9, index%9
	if board[row][col] != '.' {
		return dfs(board, index+1)
	}
	check := func(num byte) bool {
		for i:=0;i<9;i++ {
			if board[row][i] == num || board[i][col] == num || board[row/3*3+i/3][col/3*3+i%3] == num {
				return true
			}
		}
		return false
	}
	for num := byte('1'); num <= '9'; num++ {
		if check(num) {
			continue
		}
		board[row][col] = num
		if dfs(board, index+1) {
			return true
		}
		board[row][col] = '.'
	}
	return false
}
