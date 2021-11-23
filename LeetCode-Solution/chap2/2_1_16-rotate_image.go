/*
You are given an n × n 2D matrix representing an image.
Rotate the image by 90 degrees (clockwise).
Follow up: Could you do this in-place?

* @Author: Yajun
* @Date:   2021/11/22 20:19
*/

package chap2

// time: O(n^2)
// 暴力解法：分圈处理
func rotateImage(matrix [][]int) {
	var (
		n              = len(matrix)
		tr, tc, dr, dc int
		tmp            int
	)

	tr, tc = 0, 0
	dr, dc = n-1, n-1
	for tr < dr {
		for i := 0; i < dc-tc; i++ {
			tmp = matrix[tr][tc+i]
			matrix[tr][tc+i] = matrix[dr-i][tc]
			matrix[dr-i][tc] = matrix[dr][dc-i]
			matrix[dr][dc-i] = matrix[tr+i][dc]
			matrix[tr+i][dc] = tmp
		}
		tr++
		tc++
		dr--
		dc--
	}
}

/*


D----C  		 A----D
|    |    ==>    |    |
A----B           B----C

1. 斜对角线交换
B----C
|    |
A----D

2. 水平线交换
A----D
|    |
B----C
*/

// time: O(n^2)
func rotateImageB(matrix [][]int) {
	var (
		n = len(matrix)
	)
	// 斜对角线交换
	for r := 0; r < n-1; r++ {
		for c := 0; c < n-1-r; c++ {
			matrix[r][c], matrix[n-1-c][n-1-r] = matrix[n-1-c][n-1-r], matrix[r][c]
		}
	}

	// 水平交换
	for r := 0; r < n/2; r++ {
		for c := 0; c < n; c++ {
			matrix[r][c], matrix[n-1-r][c] = matrix[n-1-r][c], matrix[r][c]
		}
	}
}
