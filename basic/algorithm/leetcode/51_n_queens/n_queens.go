package n_queens

import "strings"

func solveNQueens(n int) [][]string {
	res := [][]string{}
	track := make([]string, n)
	for i := 0; i < n; i++ {
		track[i] = strings.Repeat(".", n)
	}
	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			tmp := make([]string, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		for col := 0; col < n; col++ {
			if !isVaild(track, row, col) {
				continue
			}
			tmp := []rune(track[row])
			tmp[col] = 'Q'
			track[row] = string(tmp)
			backtrack(row + 1)
			tmp = []rune(track[row])
			tmp[col] = '.'
			track[row] = string(tmp)
		}
	}
	backtrack(0)
	return res
}

func isVaild(s []string, row, col int) bool {
	// 同列
	for i := row - 1; i >= 0; i-- {
		if s[i][col] == 'Q' {
			return false
		}
	}
	// 左上角
	for i := 1; row-i >= 0 && col-i >= 0; i++ {
		if s[row-i][col-i] == 'Q' {
			return false
		}
	}
	// 右上角
	for i := 1; row-i >= 0 && col+i < len(s[0]); i++ {
		if s[row-i][col+i] == 'Q' {
			return false
		}
	}
	return true
}
