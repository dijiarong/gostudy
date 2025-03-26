package solve_sudoku

func solveSudoku(board [][]byte) {
	found := false
	var backtrack func(index int)
	backtrack = func(index int) {
		if found {
			return
		}
		if index == 81 {
			found = true
			return
		}
		row, col := index/9, index%9
		if board[row][col] != '.' {
			backtrack(index + 1)
			return
		}
		for i := byte('1'); i <= '9'; i++ {
			// 剪枝
			if !isVaild(board, row, col, i) {
				continue
			}
			board[row][col] = i
			backtrack(index + 1)
			if found {
				return
			}
			board[row][col] = '.'
		}
	}
	backtrack(0)
}

func isVaild(board [][]byte, row, col int, val byte) bool {
	// 同行同列不能出现重复的
	for i := 0; i < 9; i++ {
		if board[row][i] == val || board[i][col] == val {
			return false
		}
	}
	// 同一个3X3也不允许出现
	for i := 0; i < 9; i++ {
		if board[row/3+i/3][col/3+i%3] == val {
			return false
		}
	}
	return true
}
