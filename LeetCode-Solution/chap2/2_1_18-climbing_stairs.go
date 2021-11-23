/*
You are climbing a stair case. It takes n steps to reach to the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

* @Author: Yajun
* @Date:   2021/11/22 23:35
*/

package chap2

/*
f(n) = f(n-1) + f(n-2)
*/

// time: O(2^n); space: O(2^n)
func climbingStairs(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	return climbingStairs(n-1) + climbingStairs(n-2)
}

// f(n) = f(n-1) + f(n-2)
func climbingStairsB(n int) int {
	var (
		f1, f2 int
	)
	if n == 0 || n == 1 {
		return 1
	}
	f1, f2 = 1, 1
	for i := 2; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}
	return f2
}
