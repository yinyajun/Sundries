package main

import (
	"fmt"
	"math"
)

/*
https://cp-algorithms.com/data_structures/sparse-table.html

任何一个非负整数，都可以表达成数个2的幂次方之和
本质上，就是2进制表示

同样的，一段区间也可以表示成数个子区间之和，每个子区间的长度为2的幂次方
sparse table预计算所有长度为2的幂次方的子区间答案
将query拆解到数个子区间，找到预计算的答案，然后再拼接成最终的答案

st[i][j]: store the answer for the range [i, i + 2^j - 1] of length 2^j

st[N][K]
2^K <= N ==》  K <= log2N

[i, i + 2^j -1] 可以分为两个子区间[i, i +2^(j-1) - 1], [i + 2^(j-1), i+2^j-1]

st[i][j] = f( st[i][j-1], st[ i + (i << j-1) ][j-1])

base:
st[i][0] : [i, i] = arr[i]

*/

type RMQ struct {
	st [][]int
}

func NewRMQ(arr []int) *RMQ {
	n := len(arr)
	k := int(math.Log2(float64(n)))
	r := &RMQ{st: make([][]int, n)}
	for i := range r.st {
		r.st[i] = make([]int, k)
		r.st[i][0] = arr[i]
	}

	// [i, i+2^j-1], i+2^j-1 <N ==》 i+2^j <=N
	for j := 1; j < k; j++ {
		for i := 0; i+(1<<(j-1)) < n; i++ {
			r.st[i][j] = r.st[i][j-1] + r.st[i+(1<<(j-1))][j-1]
		}
	}
	return r
}

func (q *RMQ) Query(l, r int) int {
	var sum int
	k := int(math.Log2(float64(len(q.st))))
	for j := k; j >= 0; j-- {
		if (1 << j) <= (r - l + 1) {
			sum += q.st[l][j]
			l += 1 << j
		}
	}
	return sum
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	q := NewRMQ(arr)
	fmt.Println(q.Query(3, 6))
}
