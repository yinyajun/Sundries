package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// 最长递增子序列

// 最长上升子序列经典的动态规划题目
// dp[i]: 以arr[i]结尾情况下，arr[...i]的最长上升子序列长度
// 对于dp[i]的定义，为什么要多加上arr[i]结尾这个条件？不加这个条件的话，就是arr[...i]的最长上升子序列长度。能否用这个dp[i]和arr[i+1]去构成dp[i+1]呢
// 其实是行不通的，因为不知道arr[i+1]和前面的最长上升子序列的关系

// arr[i]结尾的最长上升子序列中，在arr[...i-1]中，任意比arr[i]小的数字都可以作为倒数第二个数
// 那么到底选哪个呢？就选哪个数结尾的最长上升子序列更大的那个
// dp[i] = max{dp[j]+1}  0<=j<i  arr[j] < arr[i]
// 初始： dp[i]=1， 至少包含arr[i]自己
// 时间复杂度为O(N^2)
func LAS(arr []int) []int {
	dp := make([]int, len(arr))
	// iterate and init
	for i := 0; i < len(arr); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				dp[i] = utils.MaxInt(dp[i], dp[j]+1)
			}
		}
	}
	fmt.Println(dp)
	return generatePath(dp, arr)

}

func generatePath(dp []int, arr []int) []int {
	maxDp := 0
	index := 0

	for i := 0; i < len(dp); i++ {
		if dp[i] > maxDp {
			maxDp = dp[i]
			index = i
		}
	}

	path := make([]int, maxDp)
	last := maxDp - 1
	path[last] = arr[index]
	last--

	for i := index; i >= 0; i-- {
		if arr[i] < arr[index] && dp[i] == dp[index]-1 {
			path[last] = arr[i]
			last--
			index = i
		}
	}
	return path
}

// 使用二分查找优化计算dp数组的时间复杂度，降为O(NlogN)
// 怎么用到二分查找呢？首先外循环遍历arr，这个动不了
// 那么只能对内循环下手了
// 首先考察内循环，发现类似一维搜索，在arr[...i-1]中找到最合适的倒数第二个数，由于没有专门的数据结构来记录，只能用一维搜索的方式
// 这样就很浪费时间了，应该使用一个专门的数据结构，使得以低于线性时间的方式来更新dp[i]
// 所以需要用一个数组来记录，这里使用ends数组，key为已有的递增子序列长度-1，val为对应子序列长度中最小的结尾元素值
// 这样的话，每次遍历新元素e的时候，
// 先去ends中查找是否没有大于e的，没有的话，当前最长的递增子序列+e可以组成更长的递增子序列，所以扩展ends数组（由于添加的e大于ends中的所有元素，保证了ends数组的单调性）
// 否则，ends中找到第一个>=e的key，说明这样的一个事实
// 首先key长度递增子序列最后一个值都比e小，可以和e组成key+1长度的递增子序列
// key+1长度的递增子序列的最后一个值需要更新为更小的e
func LAS2(arr []int) []int {
	dp := make([]int, len(arr))
	ends := make([]int, len(arr))
	maxLen := 0
	ends[0] = arr[0]

	// 在ends[left. right]中找到第一个大于等于target的index
	binarySearch := func(ends []int, left, right, target int) int {
		for left <= right {
			mid := left + (right-left)/2
			if target > ends[mid] { // [mid+1, right]
				left = mid + 1
			} else if target < ends[mid] { // [left, mid-1]
				right = mid - 1
			} else { // target == ends[mid]
				right = mid - 1 // [left, mid-1] 继续向左搜索
			}
		}
		// left == right + 1
		return left // left 取值范围为[0, right+1]，存在越界的情况，不过这里正好是我们所需要的扩张时机
	}

	for i := 0; i < len(arr); i++ {
		res := binarySearch(ends, 0, maxLen, arr[i])
		maxLen = utils.MaxInt(maxLen, res)
		ends[res] = arr[i] // res<maxLen，更新已有ends；否则添加新的元素来扩大ends
		dp[i] = res + 1
		//if  res == maxLen+1 {
		//	maxLen += 1
		//	ends[res] = arr[i]
		//	dp[i] = res+1
		//} else {
		//	ends[res] = arr[i]
		//	dp[i] = res + 1
		//}
	}
	fmt.Println(dp)
	return generatePath(dp, arr)
}

func main() {
	arr := []int{2, 1, 5, 3, 6, 4, 8, 9, 7}
	fmt.Println(LAS(arr))
	fmt.Println(LAS2(arr))
}
