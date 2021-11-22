/*
The set [1,2,3,...,n] contains a total of n! unique permutations.
By listing and labeling all of the permutations in order, We get the following sequence (ie, for n = 3):
"123"
"132"
"213"
"231"
"312"
"321"
Given n and k, return the kth permutation sequence.
Note: Given n will be between 1 and 9 inclusive.

* @Author: Yajun
* @Date:   2021/11/21 16:31
*/

package chap2

import (
	"container/list"
	"solution/utils"
	"strconv"
	"strings"
)

// brute force: call k-1 nextPermutation
// time: O(n*k); space: O(1)
func permutationSequence(n, k int) string {
	utils.Assert(n > 0 && k >= 0)

	init := make([]int, n)
	for i := 0; i < n; i++ {
		init[i] = i + 1
	}
	for i := 0; i < k-1; i++ {
		nextPermutation(init)
	}
	return intArray2String(init)
}

func intArray2String(nums []int) string {
	st := strings.Builder{}
	for i := 0; i < len(nums); i++ {
		st.WriteString(strconv.Itoa(nums[i]))
		if i < len(nums)-1 {
			st.WriteString(",")
		}
	}
	return st.String()
}

// time: O(n); space: O(n)
// 使用辗转相除法，求得康拓展开的系数，记为所求的排列
func permutationSequenceB(n, k int) string {
	utils.Assert(n > 0 && k >= 0)

	var (
		res  = strings.Builder{}
		seq  = new(list.List)
		base int
		cur  *list.Element
	)

	for i := 1; i <= n; i++ {
		seq.PushBack(i)
	}

	k = k - 1 // k的计数从0开始
	base = factorial(n - 1)

	for i := n - 1; i > 0; i-- {
		// 在deque中寻找第(k / base)个数
		cur = seq.Front()
		for i := 0; i < k/base; i++ {
			cur = cur.Next()
		}
		res.WriteString(strconv.Itoa(cur.Value.(int)))
		seq.Remove(cur) // ！移除已经选择的元素

		k = k % base
		base /= i
	}
	res.WriteString(strconv.Itoa(seq.Front().Value.(int)))
	return res.String()
}

func factorial(n int) int {
	utils.Assert(n >= 0)
	if n == 0 {
		return 0
	}
	var res = 1
	for n > 1 {
		res *= n
		n--
	}
	return res
}
