/*
There are N children standing in a line. Each child is assigned a rating value.
You are giving candies to these children subjected to the following requirements:
• Each child must have at least one candy.
• Children with a higher rating get more candies than their neighbors.
What is the minimum candies you must give?

* @Author: Yajun
* @Date:   2021/11/26 19:50
*/

package chap2

import (
	"solution/utils"
)

// time: O(n); space: O(n)
func candiesB(ratings []int) int {
	var (
		n          = len(ratings)
		increments = make([]int, n)
		inc        int
		res        = n
	)

	inc = 0 // 这里的inc设置错了好多次，todo：有时间重新再来一遍
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			inc++
			increments[i] = utils.MaxInt(inc, increments[i])
		} else {
			inc = 0
		}
	}

	inc = 0
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			inc++
			increments[i] = utils.MaxInt(inc, increments[i])
		} else {
			inc = 0
		}
	}
	for i := 0; i < n; i++ {
		res += increments[i]
	}
	return res
}

// time: O(n^2); space: O(n^2)
func candiesC(ratings []int) int {
	var (
		n   = len(ratings)
		res int
	)

	for i := 0; i < n; i++ {
		d := candiesRecursive(ratings, i)
		res += d
	}
	return res
}

// if ratings[i] > ratings[i-1]: f(i) = f(i-1) + 1
// if ratings[i] < ratings[i+1]: f(i) = f(i+1) + 1
// == : f(i) = 0
// base: f(0) = f(n-1) = 1
func candiesRecursive(ratings []int, i int) int {
	if i == 0 || i == len(ratings)-1 {
		return 1
	}
	if ratings[i] > ratings[i-1] {
		return candiesRecursive(ratings, i-1) + 1
	} else if ratings[i] > ratings[i+1] {
		return candiesRecursive(ratings, i+1) + 1
	} else { // ==
		return 1
	}
}

func candiesD(ratings []int) int {
	var (
		n    = len(ratings)
		memo = make([]int, n)
		res  int
	)

	for i := 0; i < n; i++ {
		res += candiesRecursive2(ratings, memo, i)
	}
	return res
}

func candiesRecursive2(ratings []int, memo []int, i int) int {
	if memo[i] == 0 {
		memo[i] = 1
	}
	if i > 0 && ratings[i] > ratings[i-1] {
		memo[i] = candiesRecursive2(ratings, memo, i-1) + 1
	}
	if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
		memo[i] = candiesRecursive2(ratings, memo, i+1) + 1
	}
	return memo[i]
}

// time: O(n); space: O(1)
func candies(ratings []int) int {
	var (
		i            = nextMin(ratings, 0)
		res          = regionSum(0, i)
		lSize, rSize int
		next         int
	)
	i++
	lSize = 1
	for i < len(ratings) {
		if ratings[i] > ratings[i-1] { // strict up
			lSize++
			res += lSize
			i++
		} else if ratings[i] < ratings[i-1] { // strict down
			next = nextMin(ratings, i-1)
			rSize = next - (i - 1) + 1
			res += regionSum(i-1, next)
			if lSize > rSize {
				res -= rSize
			} else {
				res -= lSize
			}
			i = next + 1
			lSize = 1
		} else { // equal, either up or down is ok
			lSize = 1
			res += 1
			i++
		}
	}
	return res
}

// 从ratings[start...]中寻找下降沿（严格下降）结束的位置（也就是上升沿开始的位置）
func nextMin(ratings []int, start int) int {
	for i := start; i < len(ratings)-1; i++ {
		if ratings[i] > ratings[i+1] {
			continue
		}
		// ratings[i] <= ratings[i+1]
		return i
	}
	return len(ratings) - 1
}

func regionSum(left, right int) int {
	n := right - left + 1
	return n * (n + 1) / 2
}

func candies2(ratings []int) int {
	var (
		i            = nextMin2(ratings, 0)
		res, _       = regionSumAndSize(ratings, 0, i)
		lSize, rSize int
		next         int
		same         int
		candies      int
	)
	i++
	same = 1
	lSize = 1

	for i < len(ratings) {
		if ratings[i] > ratings[i-1] { // strict up
			lSize++
			res += lSize
			same = 1
			i++
		} else if ratings[i] < ratings[i-1] { // instruct down
			next = nextMin2(ratings, i-1)
			candies, rSize = regionSumAndSize(ratings, i-1, next)
			res += candies
			if lSize >= rSize {
				res -= rSize
			} else {
				res -= same*lSize - rSize + same*rSize // 移除左右，单独加上右边
			}
			same = 1
			i = next + 1
			lSize = 1
		} else {
			res += lSize
			same++
			i++
		}
	}
	return res
}

// 从ratings[start...]中寻找下降沿（非严格下降）结束的位置（也就是上升沿开始的位置）
func nextMin2(ratings []int, start int) int {
	for i := start; i < len(ratings)-1; i++ {
		if ratings[i] >= ratings[i+1] {
			continue
		}
		// ratings[i] < ratings[i+1]
		return i
	}
	return len(ratings) - 1
}

func regionSumAndSize(ratings []int, left, right int) (int, int) {
	var res, size int
	res = 1
	size = 1
	for i := right - 1; i >= left; i-- {
		if ratings[i] > ratings[i+1] {
			size++
			res += size
		} else { // ==
			res += size
		}
	}
	return res, size
}
