/*
There are two sorted arrays A and B of size m and n respectively. Find the median of the two sorted
arrays. The overall run time complexity should be O(log(m + n)).
*
* @Author: Yajun
* @Date:   2021/10/5 11:18
*/

package chap2

import "solution/utils"

// time: O(m+n); space: O(m+n)
func findMedianSortedArrays(a, b []int) float32 {
	total := len(a) + len(b)
	aux := make([]int, total/2+1)

	// merge函数，找第k个（start from 1），需要数组长度为k，[0, k-1]
	findKth := func(a, b, aux []int, k int) int {
		var i, j int // a, b的指针
		for m := 0; m < k; m++ {
			if i < len(a) && j < len(b) {
				if a[i] < b[j] {
					aux[m] = a[i]
					i++
				} else {
					aux[m] = b[j]
					j++
				}
			}
			if i >= len(a) {
				aux[m] = b[j]
				j++
			}
			if j >= len(b) {
				aux[m] = a[i]
				i++
			}
		}
		return aux[k-1]
	}

	if total&1 == 1 { // odd
		return float32(findKth(a, b, aux, total/2+1))
	} else { // even
		return float32(findKth(a, b, aux, total/2)+findKth(a, b, aux, total/2+1)) / 2
	}
}

// time: O(m+n); space: O(m+n)
func findMedianSortedArraysB(a, b []int) float32 {
	total := len(a) + len(b)
	aux := make([]int, total/2+1)

	// merge函数，找第k个（start from 1），需要数组长度为k，[0, k-1]
	findKth := func(a, b, aux []int, k int) int {
		var i, j, m int // a, b, aux的指针
		for ; m < k; m++ {
			if i > len(a) {
				aux[m] = b[j]
				j++
			} else if j > len(b) {
				aux[m] = a[i]
				i++
			} else if a[i] < b[j] {
				aux[m] = a[i]
				i++
			} else {
				aux[m] = b[j]
				j++
			}
		}
		return aux[k-1]
	}

	if total&1 == 1 { // odd
		return float32(findKth(a, b, aux, total/2+1))
	} else { // even
		return float32(findKth(a, b, aux, total/2)+findKth(a, b, aux, total/2+1)) / 2
	}
}

// time: O(log(m+n)); space: O(log(m+n))
// 上面的merge做法，得到的aux已经有序，而求k-th操作，根本不需要有序那么复杂，所以时间复杂度应该还可以降下去
// 基于：每次额能删除k-th元素之前的的一个元素，那么进行k次，即可找到k-th，如果每次能删除一半元素呢，使用二分查找。
// 1. a[k/2-1] == b[k/2-1] => 找到第k大的元素， return a[k/2-1]
// 2. a[k/2-1] < b[k/2-1] => [a[0], a[k/2-1]]肯定在topK中, 删除之，然后从a[k/2...]和b中寻找k-k/2
// 3. a[k/2-1] > b[k/2-1] => [b[0], b[k/2-1]]肯定在topK中, 删除之，然后从a和b[k/2...]中寻找k-k/2
// 递归过程中，不断缩小的量为a的搜索区间，b的搜索区间，k，所以base情况针对这些量
// base:
// 1. k==1, return min(a[0], b[0])
// 2. a搜索区间==0, return b[k-1]
// 3. b搜索区间==0 , return a[k-1]
func findMedianSortedArraysC(a, b []int) float32 {
	total := len(a) + len(b)
	if total&1 == 1 { // odd
		return float32(findKth1(a, len(a), b, len(b), total/2+1))
	} else { // even
		return float32(findKth1(a, len(a), b, len(b), total/2)+findKth1(a, len(a), b, len(b), total/2+1)) / 2
	}
}

// 在a[0,m)和b[0,n)中搜索第k个元素
func findKth1(a []int, m int, b []int, n int, k int) int {
	if m == 0 {
		return b[k-1]
	}
	if n == 0 {
		return a[k-1]
	}
	if k == 1 { // 然后再放置该base条件，确保a[0], b[0]存在
		return utils.MinInt(a[0], b[0])
	}
	// divide k into two parts
	ia := utils.MinInt(k/2, m) // 防止a[k/2-1]越界
	ib := k - ia

	if a[ia-1] < b[ib-1] {
		// remove a[:ia-1], a becomes a[ia, m)
		return findKth1(a[ia:], m-ia, b, ib, k-ia)
	} else if a[ia-1] > b[ib-1] {
		// remove b[:ib-1], b becomes b[ib, n)
		return findKth1(a, m, b[ib:], n-ib, k-ib)
	} else { // a[ia-1] == b[ib-1]
		return a[ia-1]
	}
}

// time: O(log(m+n)); space: O(log(m+n))
// 同上，简单的改个数组的表达方式，不创建新的slice
func findMedianSortedArraysD(a, b []int) float32 {
	total := len(a) + len(b)

	if total&1 == 1 { // odd
		return float32(findKth2(a, 0, len(a)-1, b, 0, len(b)-1, total/2+1))
	} else { // even
		return float32(findKth2(a, 0, len(a)-1, b, 0, len(b)-1, total/2)+
			findKth2(a, 0, len(a)-1, b, 0, len(b)-1, total/2+1)) / 2
	}
}

// a[aLo, aHi]和b[bLo, bHi]中寻找第k个元素
func findKth2(a []int, aLo, aHi int, b []int, bLo, bHi int, k int) int {
	if aLo > aHi {
		return b[bLo+k-1]
	}
	if bLo > bHi {
		return a[aLo+k-1]
	}
	if k == 1 {
		return utils.MinInt(a[aLo], b[bLo])
	}
	// divide k into two parts
	ia := utils.MinInt(k/2, aHi-aLo+1)
	ib := k - ia
	if a[aLo+ia-1] < b[bLo+ib-1] { // remove a[aLo: aLo+ia-1]
		return findKth2(a, aLo+ia, aHi, b, bLo, bHi, k-ia)
	} else if a[aLo+ia-1] > b[bLo+ib-1] {
		return findKth2(a, aLo, aHi, b, bLo+ib, bHi, k-ib)
	} else {
		return a[aLo+ia-1]
	}
}
