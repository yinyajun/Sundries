package main

import "fmt"

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
// 由于s不单调，需要构造辅助数组，用来寻找sum[0...i] >= s[j] -k的最早出现的位置
// 而辅助数组构造方法很简单
// 例如s=[0,1,3,2,7,5], help=[0,1,3,3,7,7]
// 因为只关心>=某个值最早出现的位置。
// 这里的思路和单调栈差不多，在单调栈中，将不可能的答案直接删除，不保存在数据结构中，而答案在数据结构中，可以用o(1)时间获得
// 而这里将不可能的答案处理掉（m>n, s[m]>s[n], 由于查找大于等于某数k的最早位置，如果s[n]>=k，那么s[m]>=k且m<n）,当然这里并没有删除
// 这样辅助数组help就是单调的，可以使用二分查找

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

func preprocess(a []int) []int {
	cumSum := make([]int, len(a)+1)
	cumSum[0] = 0
	for i := 0; i < len(a); i++ {
		cumSum[i+1] = cumSum[i] + a[i]
	}
	return cumSum
}

func preprocess2(a []int) []int {
	cumSum := []int{}
	cumSum = append(cumSum, 0)
	for i := 0; i < len(a); i++ {
		cumSum = append(cumSum, cumSum[len(cumSum)-1]+a[i])
	}
	return cumSum
}

func preprocess3(a []int) []int {
	cumSum := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		if i == 0 {
			cumSum[i] = a[i]
			continue
		}
		cumSum[i] = cumSum[i-1] + a[i]
	}
	return cumSum
}

// a[i...j]
func sumRange(cumSum []int, i, j int) int {
	// 省略非法检测
	if i == 0 {
		return cumSum[j]
	}
	return cumSum[j] - cumSum[i-1]
}

func sumRange1(cumSum []int, i, j int) int {
	return cumSum[j+1] - cumSum[i]
}

func preprocess4(a []int) []int {
	cumSum := make([]int, len(a)+1)
	cumSum[0] = 0
	for i := 1; i <= len(a); i++ {
		cumSum[i] = cumSum[i-1] + a[i-1]
	}
	return cumSum
}

//func main() {
//	//b := []int{-100, 1, 2, 3, -2, 3, -6, 0, 1, 0}
//	//fmt.Println(MaxLengthUnderK(b, -10))
//	a := []int{2, 6, 7, 4}
//	r := preprocess(a)
//	fmt.Println(r)
//	r = preprocess2(a)
//	fmt.Println(r)
//	r = preprocess4(a)
//	fmt.Println(r)
//	fmt.Println(sumRange1(r, 0,3))
//	r = preprocess3(a)
//	fmt.Println(r)
//	fmt.Println(sumRange(r, 0, 3))
//}

func main() {
	num := []int{2, 5, 7, 9, 10}
	fmt.Println(_find(num, 0, len(num)-1, 11))
}
