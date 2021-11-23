/*
The gray code is a binary numeral system where two successive values differ in only one bit.
Given a non-negative integer n representing the total number of bits in the code, print the sequence of
gray code. A gray code sequence must begin with 0.
For example, given n = 2, return [0,1,3,2]. Its gray code sequence is:
00 - 0
01 - 1
11 - 3
10 - 2
Note:
• For a given n, a gray code sequence is not uniquely defined.
• For example, [0,2,3,1] is also a valid gray code sequence according to the above definition.
• For now, the judge is able to judge based on one instance of gray code sequence. Sorry about that.

* @Author: Yajun
* @Date:   2021/11/22 23:53
*/

package chap2

/*
自然二进制码 转 格雷码
g_0 = b_0 (最高位保持不变，作为基准值)
g_i = b_i ^ b_{i-1}

intuitive:
自然二进制中连续的两个数
1. b0 b1 b2 b3
2. b0 b1 b2 b3 + 1

	（假如个位上为0， 变为 b0 b1 b2 ~b3)
	转为gray码
	1. b0 b0^b1 b1^b2 b2^b3
	2. b0 b0^b1 b1^b2 b2^(~b3)
	可以发现gray码中仅有一个bit不同

	（假如个位上为1，十位为0， 变为 b0 b1 ~b2 ~b3)
	转为gray码
	1. b0 b0^b1 b1^b2 b2^b3
	2. b0 b0^b1 b1^(~b2) (~b2)^(~b3)
	可以发现gray码中仅有一个bit不同

格雷码 转 自然二进制码
b_0 = g_ 0
b_i = g_i ^ g_{i-1}

*/

// time: O(2^n); space: O(1)
func grayCode(n int) []int {
	var (
		size = 1 << n
		res  = make([]int, size)
	)

	for i := 0; i < size; i++ {
		res[i] = i ^ (i >> 1)
	}
	return res
}

// time: O(2^n); space: O(1)
// 递归求解 f(n) = [f(n-1)..., 1<<(n-1) + reverse(f(n-1))...]
func grayCodeB(n int) []int {
	res := make([]int, 1<<n)
	_recursiveB(res, n)
	return res
}

func _recursiveB(res []int, bit int) int {
	if bit == 0 {
		res[0] = 0
		return 1
	}
	before := _recursiveB(res, bit-1)
	for idx, n := range res[:before] {
		base := 1 << (bit - 1)
		//res[idx] = n
		res[base<<1-1-idx] = n | base // 对称（|这里等价于+）
	}
	return 1 << bit
}

// time: O(2^n); space: O(1)
// 将上述递归转为非递归
func grayCodeC(n int) []int {
	res := make([]int, 0)
	res = append(res, 0)

	for i := 0; i < n; i++ {
		base := 1 << i
		for j := len(res) - 1; j >= 0; j-- { // 对称
			res = append(res, res[j]|base)
		}
	}
	return res
}
