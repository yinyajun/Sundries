package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// 给定一个无序数组，元素可正可负可0，求得所有子数组中，累加和小于等于k的最长子数组

// 这里和上一题求区间和=k的最长子数组又有点不同
// 上一题是等号，可以作为map中的key，通过O(1)的复杂度查找
// 这一题是不等号，用map并不能提高查找的速度。如果能构造出单调情况，可以使用二分查找来降低查找的复杂度到O(logN)

// 思路：find i, min_i s(j)-s(i)<= k
// <==> min_i s(i) >= s(j) - k
// s(i)递增，在递增的数组中寻找大于某个值的最小值
// 时间复杂度O(NlogN), 空间复杂度O(N)

// @@@@@@ 上述思路是错误的，s(i)并不是递增的，因为数组元素有正有负。
func MaxLengthUnderK(a []int, k int) int {
	sum := make([]int, len(a))
	length := 0

	for j := 0; j < len(a); j++ {
		// 构造前缀和数组
		if j == 0 {
			sum[j] = a[j]
		} else {
			sum[j] = sum[j-1] + a[j]
		}
		// 通过二分查找找到符合条件的左边界
		i := _find2(sum, 0, j, sum[j]-k)
		if j-i > length {
			length = j - i
		}
	}
	return length
}

// 如何构造出递增数组？仍然是这个思路
// 思路：find i, min_i s(j)-s(i)<= k
// <==> min_i s(i) >= s(j) - k
// 由于s不单调，需要构造辅助数组，用来寻找sum[0...i] >= s[j] - k 的最早出现的位置
// 而辅助数组构造方法很简单
// 例如s=[0,1,3,2,7,5], help=[0,1,3,3,7,7]
// 因为只关心>=某个值最早出现的位置。
// 这里的思路和单调栈差不多，在单调栈中，将不可能的答案直接删除，不保存在数据结构中，而答案在数据结构中，可以用o(1)时间获得
// 而这里将不可能的答案处理掉（m>n, s[m]>s[n], 由于查找大于等于某数k的最早位置，如果s[n]>=k，那么s[m]>=k且m<n）,当然这里并没有删除
// 这样辅助数组help就是单调的，可以使用二分查找
// s[i+1] = sum(a[0...i])
// s[i+1] = s[i] + a[i]
func MaxLengthUnderK2(a []int, k int) int {
	var length, sum int
	cum := make([]int, len(a)+1)
	cum[0] = sum

	for i := 0; i < len(a); i++ {
		sum += a[i]
		cum[i+1] = utils.MaxInt(cum[i], sum)
		idx := _find2(cum, 0, i+1, sum-k)
		fmt.Println("target:", sum-k, "array:", cum[:i+2], "left bound:", idx, "right bound:", i+1, "len:", i+1-idx)
		if i-idx+1 > length {

			length = i - idx + 1
		}
	}
	return length
}

func main() {
	num := []int{2, 5, 7, 9, 10}
	fmt.Println(_find(num, 0, len(num)-1, 11))
	a := []int{3, -2, -4, 0, 6}
	fmt.Println(MaxLengthUnderK2(a, -2))
}

// 在递增数组中，寻找大于target的最小索引
// 换种说法，找到第一个大于target的索引
func _find(sum []int, lo, hi int, target int) int {
	var mid int
	// [lo, hi]
	for lo <= hi {
		mid = lo + (hi-lo)/2
		if sum[mid] == target {
			hi = mid - 1
		} else if sum[mid] < target {
			lo = mid + 1
		} else { // sum[mid] > target
			hi = mid - 1
		}
	}
	return lo
}

func _find2(sum []int, lo, hi, target int) int {
	var mid int
	for lo <= hi {
		mid = lo + (hi-lo)/2
		if sum[mid] >= target { // left part
			hi = mid - 1
		} else { // sum[mid] < target
			lo = mid + 1 // right part
		}
	}
	// when lo == hi == mid and target <= sum[mid], left part, hi-- , so lo > target, return lo
	// when lo == hi == mid and target > sum[mid], right part, lo++ , so lo > target, return lo
	return lo
}
