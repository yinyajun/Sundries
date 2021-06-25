package main

import "fmt"

func SpatialOrderPrint(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	tR, tC := 0, 0
	dR, dC := len(matrix)-1, len(matrix[0])-1
	for (tR <= dR) && (tC <= dC) {
		fmt.Println(tR, tC, dR, dC)
		printEdge(matrix, tR, tC, dR, dC)
		tR++
		tC++
		dR--
		dC--
	}
}

// 注意：二维数组中，别用xy，因为是反的
func printEdge(matrix [][]int, tr, tc, dr, dc int) {
	if tr == dr { // 只有一行
		for col := tc; col <= dc; col++ {
			fmt.Println(matrix[tr][col])
		}
		return
	}

	if tc == dc { // 只有一列
		for row := tr; row <= dr; row++ {
			fmt.Println(matrix[row][tc])
		}
		return
	}

	// 正常情况
	row, col := tr, tc
	for col < dc {
		fmt.Println(matrix[row][col])
		col++
	}

	for row < dr {
		fmt.Println(matrix[row][col])
		row++
	}

	for col > tc {
		fmt.Println(matrix[row][col])
		col--
	}

	for row > tr {
		fmt.Println(matrix[row][col])
		row--
	}

}

//func main() {
//	array := [][]int{{1, 2, 3, 4, 5},{29, 9324,54,534,75}, {6,7,8,9,0}}
//	SpatialOrderPrint(array)
//}
