package main

// 子矩阵和子数组不一样，因为矩阵的区域和更加复杂：cum[x1+1][y1+1] - cum[x0][y1+1] - cum[x1+1][y0] + cum[x0][y0]
// 这样寻找合适的左上角将变得更加复杂

// 前缀矩阵的求解也需要需要O(N^2)的时间，此时遍历去找一个合适的左上角元素需要O(N^2)时间，那么总时间将会是O(N^4)

// 期望的复杂度是O(N^2)，但是实在没有办法
// 回顾上一题，可以做到O(N)时间内找到数组中最大的子数组
// 数组可以认为是1*N的矩阵
// 将两行数组合并成一行，就可以O(N)时间求得2*N的矩阵中最大的行数为2的子矩阵

// 那么，遍历一个矩阵所有的行数组合（1行的，2行的，3行的... n行的）
// 组合数为n+(n-1) +..+1= O(N^2)
// 每次组合中，花费O(N)时间计算出最大的子数组即可，最后总复杂度为O(N^3)

func MaxSumSubMatrix(mat [][]int) int {
	var nums []int
	var ans int
	var max int
	//for n := 1; n <= len(mat); n++ {
	//	for j := 0; j+n <= len(mat); j++ { // [j, j+n)
	//		fmt.Println(j, j+n)
	//	}
	//}

	// 每次外循环都确定了一个矩阵行数，这样计算矩阵合并，不能利用之前的结果，会有重复计算
	//for n := 0; n < len(mat); n++ {
	//	for j := 0; j+n < len(mat); j++ { // [j, j+n]
	//	}
	//}
	if len(mat) <= 0 {
		return 0
	}
	// 外循环为起始行数，内循环为合并矩阵数，这样下一次合并可以复用上次合并的数组
	for i := 0; i < len(mat); i++ {
		nums = make([]int, len(mat[0]))
		for n := 0; i+n < len(mat); n++ { // [i,i+n]
			for k := 0; k < len(mat[0]); k++ {
				nums[k] += mat[i+n][k]
			}
			max = maxSumSubarray(nums)
			if max > ans {
				ans = max
			}
		}
	}
	return ans
}

func maxSumSubarray(nums []int) int {
	var ans int
	var minVal int // sum[0,-1] = 0
	var cum int

	for i := 0; i < len(nums); i++ {
		cum += nums[i]
		if cum-minVal > ans {
			ans = cum - minVal
		}
		if cum < minVal {
			minVal = cum
		}
	}
	return ans
}

//func main() {
//	mat := [][]int{
//		{-1, -1, -1},
//		{-1, 2, 2},
//		{-1, -1, -1},
//	}
//	fmt.Println(MaxSumSubMatrix(mat))
//}
