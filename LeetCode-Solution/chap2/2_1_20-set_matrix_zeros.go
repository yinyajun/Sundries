/*
Given a m × n matrix, if an element is 0, set its entire row and column to 0. Do it in place.
Follow up: Did you use extra space?
A straight forward solution using O(mn) space is probably a bad idea.
A simple improvement uses O(m + n) space, but still not the best solution.
Could you devise a constant space solution?

* @Author: Yajun
* @Date:   2021/11/23 14:21
*/

package chap2

// time: O(m*n); space: O(m + n)
func setZeros(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	row := make([]bool, m)
	column := make([]bool, n)

	for r := range matrix {
		for c := range matrix[r] {
			if matrix[r][c] == 0 {
				row[r] = true
				column[c] = true
			}
		}
	}

	for i := 0; i < m; i++ {
		if row[i] {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 0; j < n; j++ {
		if column[j] {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
}

// time: O(m*n); space: O(1)
// 复用第一行和第一列
func setZerosB(matrix [][]int) {
	var (
		firstRowZero, firstColZero bool
		m, n                       = len(matrix), len(matrix[0])
	)

	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstRowZero = true
			break
		}
	}

	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
			break
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstRowZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	if firstColZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
