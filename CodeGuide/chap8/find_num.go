package main

// 行列中都排好序的矩阵中寻找一个数 【O(N+M)时间，O(1)空间】

func findNum(matrix [][]int, num int) bool {
	if len(matrix) == 0 {
		return false
	}
	row, col := len(matrix), len(matrix[0])
	if col == 0 {
		return false
	}

	// [0,i)<=num
	var i, j int
	for ; i < row && matrix[i][0] <= num; i++ {
	} // i == row || matrix[i][0] >num

	i--
	for ; j < col && matrix[i][j] <= num; j++ {
	} // j== col || matrix[i][j] > num
	j--
	return matrix[i][j] == num
}

// 从右上角开始
func findNum2(matrix [][]int, num int) bool {
	if len(matrix) == 0 {
		return false
	}
	row, col := len(matrix), len(matrix[0])
	if col == 0 {
		return false
	}

	r, c := 0, col-1
	for r < row && c >= 0 {
		if matrix[r][c] == num {
			return true
		} else if matrix[r][c] > num {
			c--
		} else { // matrix[r][c] < num
			r++
		}
	}
	return false
}

//func main() {
//	matrix := [][]int{
//		{0, 1, 2, 5},
//		{2, 3, 4, 7},
//		{4, 4, 4, 8},
//		{5, 7, 7, 9},
//	}
//	fmt.Println(findNum(matrix, 9))
//	fmt.Println(findNum2(matrix, 9))
//}
