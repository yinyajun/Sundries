package main

// 分圈处理
func rotateMatrix(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	tr, tc := 0, 0
	dr, dc := len(matrix)-1, len(matrix[0])-1

	for tr <= dr && tc <= dc {
		rotate90(matrix, tr, tc, dr, dc)
		tr++
		dr--
		tc++
		dc--
	}
}

// matrix是N*N的，没有特殊情况
func rotate90(matrix [][]int, tr, tc, dr, dc int) {
	for i := 0; i < (dc - tc); i++ {
		tmp := matrix[tr][tc+i]
		matrix[tr][tc+i] = matrix[dr-i][tc]
		matrix[dr-i][tc] = matrix[dr][dc-i]
		matrix[dr][dc-i] = matrix[tr+i][dc]
		matrix[tr+i][dc] = tmp
	}
}

//func main() {
//	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
//	rotateMatrix(matrix)
//	fmt.Println(matrix)
//}
