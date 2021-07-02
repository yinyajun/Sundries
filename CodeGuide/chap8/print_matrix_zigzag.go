package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// (0,0)
// (0,1),(1,0)
// (2,0),(1,1),(0,2)
// (0,3),(1,2),(2,1)
// (2,2),(1,3)
func PrintMatrixZigZag(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	row, col := len(matrix), len(matrix[0])
	rowFirst := true

	var r, c int
	for s := 0; s <= row-1+col-1; s++ { // 有多少根斜线
		// 每根斜线的起点
		if rowFirst {
			r = utils.MinInt(s, row-1) // 保证不能越界
			c = s - r
			for r >= 0 && c < col {
				fmt.Print(matrix[r][c], " ")
				r--
				c++
			}
		} else {
			c = utils.MinInt(s, col-1) // 保证不能越界
			r = s - c
			for r < row && c >= 0 {
				fmt.Print(matrix[r][c], " ")
				r++
				c--
			}
		}
		rowFirst = !rowFirst
	}
}

// 上面的方法和书上的方法是等效的
// 书上的方法：维护每根斜线的两个端点
// 上端点：先向右移，然后再向下移（对应于上面rowFirst=false。col先增加，row不变；当col到达最右后，col保持不变，然后row增加）
// 下端点：先向下移，然后再想右移
func PrintMatrixZigZag2(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	row, col := len(matrix), len(matrix[0])
	fromUp := false
	tr, tc, dr, dc := 0, 0, 0, 0
	// 上端点向右，向下
	for tr <= row-1 {
		printLevel(matrix, tr, tc, dr, dc, fromUp)
		fromUp = !fromUp
		//对于上端点
		if tc == col-1 { //已经到了最右边，向下
			tr++
		} else { // 还没有到最右边，向右
			tc++
		}
		//对于下端点
		if dr == row-1 { // 已经到了最下边，向右
			dc++
		} else { // 还没有到最下边，向下
			dr++
		}
	}
}

func printLevel(m [][]int, tr, tc, dr, dc int, fromUp bool) {
	if fromUp {
		for tr <= dr {
			fmt.Print(m[tr][tc], " ")
			tr++
			tc--
		}
	} else {
		for dr >= tr {
			fmt.Print(m[dr][dc], " ")
			dr--
			dc++
		}
	}
}

//func main() {
//	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
//	PrintMatrixZigZag(matrix)
//	fmt.Println()
//	PrintMatrixZigZag2(matrix)
//}
