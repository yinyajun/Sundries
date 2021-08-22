package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// 最大由1组成的正方形的边长

// 遍历每一个元素，每个元素作为正方形的左上角，O(N^2)
// 确定一个元素作为左上角，判断能组成边长为1、2、3...N的正方形，O(N)
// 确定左上角，确定边长后，遍历4条边是否为1，O(N)
// 所以暴力解法，花费O(N^4)
func max_1_rect(mat [][]int) int {
	n := len(mat)
	var ans int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// 边长==1的情况
			if mat[i][j] == 1 {
				ans = utils.MaxInt(ans, 1)
			} else {
				continue
			}
			// 边长>1的情况
			for l := 2; l <= utils.MinInt(n-i, n-j); l++ { // 每条边只要检测l-1长度
				flag := true
				for k := 0; k < l-1; k++ {
					if mat[i][j+k] != 1 || mat[i+k][j+l-1] != 1 || mat[i+l-1][j+l-1-k] != 1 || mat[i+l-1-k][j] != 1 {
						flag = false
						break // 这里break是中止边扫描
					}
				}
				if !flag {
					continue // 该边长没有可能了
				}
				fmt.Println(i, j, l, flag)
				ans = utils.MaxInt(ans, l)
			}
		}
	}
	return ans
}

func max_1_rect2(mat [][]int) int {
	n := len(mat)
	var ans int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for l := 1; l <= utils.MinInt(n-i, n-j); l++ {
				if hasSize(mat, i, j, l) {
					ans = utils.MaxInt(ans, l)
				}
			}
		}
	}
	return ans
}

func max_1_rect3(mat [][]int) int {
	n := len(mat)
	var ans int
	right, down := setBorderMatrix(mat)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for l := 1; l <= utils.MinInt(n-i, n-j); l++ {
				if hasSize2(i, j, l, down, right) { //O(1)时间
					ans = utils.MaxInt(ans, l)
				}
			}
		}
	}
	return ans
}

func setBorderMatrix(mat [][]int) ([][]int, [][]int) {
	r, c := len(mat), len(mat[0])
	right := make([][]int, r)
	down := make([][]int, r)
	for i := 0; i < r; i++ {
		right[i] = make([]int, c)
		down[i] = make([]int, c)
	}

	// preprocess
	// 右下角
	if mat[r-1][c-1] == 1 {
		down[r-1][c-1] = 1
		right[r-1][c-1] = 1
	}

	// 最右边一列: down[i][c-1] = down[i+1][c-1] + 1
	for i := r - 2; i >= 0; i-- {
		if mat[i][c-1] == 1 {
			down[i][c-1] = down[i+1][c-1] + 1
			right[i][c-1] = 1
		}
	}

	// 最下边一行： right[r-1][j] = right[r-1][j+1] + 1
	for j := c - 2; j >= 0; j-- {
		if mat[r-1][j] == 1 {
			down[r-1][j] = 1
			right[r-1][j] = right[r-1][j+1] + 1
		}
	}

	//
	for i := r - 2; i >= 0; i-- {
		for j := c - 2; j >= 0; j-- {
			if mat[i][j] == 1 {
				down[i][j] = down[i+1][j] + 1
				right[i][j] = right[i][j+1] + 1
			}
		}
	}
	return right, down
}

func hasSize2(i, j, sz int, down, right [][]int) bool {
	if right[i][j] >= sz && down[i][j] >= sz && right[i+sz-1][j] >= sz && down[i][j+sz-1] >= sz {
		//fmt.Println(i,j,sz)
		return true
	}
	return false
}

func hasSize(mat [][]int, i, j, sz int) bool {
	if sz == 1 {
		return mat[i][j] == 1
	}
	for k := 0; k < sz-1; k++ { // 每条边只要检测l-1长度
		if mat[i][j+k] != 1 || mat[i+k][j+sz-1] != 1 || mat[i+sz-1][j+sz-1-k] != 1 || mat[i+sz-1-k][j] != 1 {
			return false
		}
	}
	return true
}

//func main() {
//	mat := [][]int{
//		{0, 1, 1, 1, 1},
//		{0, 1, 0, 0, 1},
//		{0, 1, 0, 0, 1},
//		{0, 1, 1, 1, 1},
//		{0, 1, 0, 1, 1},
//	}
//	fmt.Println(max_1_rect3(mat))
//}
