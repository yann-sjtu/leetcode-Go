package _6_valid_sudoku

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		var rows, cols, cells [9]bool
		for j := 0; j < 9; j++ {
			if num := board[i][j] - '0'; num > 0 && num <= 9 {
				if rows[num-1] {
					return false
				}
				rows[num-1] = true
			}
			if num := board[j][i] - '0'; num > 0 && num <= 9 {
				if cols[num-1] {
					return false
				}
				cols[num-1] = true
			}
			if num := board[i/3*3+j/3][i%3*3+j%3] - '0'; num > 0 && num <= 9 {
				if cells[num-1] {
					return false
				}
				cells[num-1] = true
			}
		}
	}
	return true
}

func isValidSudokuV2(board [][]byte) bool {
	var rows, cols, cells [9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]-'0'
			if num <1 || num > 9 {
				continue
			}
			if (rows[i]>>(num-1))&1 == 1 {
				return false
			}
			rows[i] ^= 1<<(num-1)
			if (cols[j]>>(num-1))&1 == 1 {
				return false
			}
			cols[j] ^= 1<<(num-1)
			if (cells[i/3*3+j/3]>>(num-1))&1 == 1 {
				return false
			}
			cells[i/3*3+j/3] ^= 1<<(num-1)
		}
	}
	return true
}
