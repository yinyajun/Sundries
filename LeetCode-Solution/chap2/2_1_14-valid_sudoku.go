/*
Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.

* @Author: Yajun
* @Date:   2021/11/21 21:43
*/

package chap2

func isValidSudoku(board [9][9]byte) bool {
	var used = [9]bool{}

	for i := 0; i < 9; i++ {
		for k := 0; k < len(used); k++ {
			used[k] = false
		}

		// check row
		for j := 0; j < 9; j++ {
			if !check(board[i][j], &used) {
				return false
			}
		}

		for k := 0; k < len(used); k++ {
			used[k] = false
		}

		// check column
		for j := 0; j < 9; j++ {
			if !check(board[j][i], &used) {
				return false
			}
		}
	}

	// check 3*3 grid
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			for k := 0; k < len(used); k++ {
				used[k] = false
			}

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if !check(board[r*3+i][c*3+j], &used) {
						return false
					}
				}
			}
		}
	}
	return true
}

func check(ch byte, used *[9]bool) bool { // 当心，golang的数组作为形参是值复制
	if ch == '.' {
		return true
	}
	if used[ch-'1'] {
		return false
	}
	(*used)[ch-'1'] = true
	return true
}
