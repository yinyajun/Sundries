package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// 计算小和，s=[1,3,5,2,4,6]
// 0+1+4+6+15=27

// 暴力解法, O(N^2)
func GetSmallSum(a []int) int {
	var ans, sum int

	for i := 1; i < len(a); i++ {
		sum = 0
		for j := 0; j < i; j++ {
			if a[j] <= a[i] {
				sum += a[j]
			}
		}
		ans += sum
	}
	return ans
}

// merge函数用O(N)时间可以知道两两之间偏序个数，但是不知道具体的偏序对（要知道具体的偏序对，需要O(N^2)）
// merge函数可以用O(N)时间理出正序关系（排序），有了正序，同时也就有了逆序。

func GetSmallSum2(a []int) int {
	return smallSum(a, 0, len(a)-1)
}

func indent(cnt int) {
	for i := 0; i < cnt; i++ {
		fmt.Print("    ")
	}
}

var cnt int

func smallSum(a []int, lo, hi int) int {
	indent(cnt)
	cnt++
	fmt.Println(lo, hi)
	if lo >= hi { // 因为left[lo, mid]而不是mid-1，所以当区间还有一个元素的时候，就该停止了
		cnt--
		indent(cnt)
		fmt.Println("==", 0)
		return 0
	}
	// lo <= hi
	mid := lo + (hi-lo)/2
	left, right := smallSum(a, lo, mid), smallSum(a, mid+1, hi)
	mergeSum := merge2(a, lo, mid, hi)
	cnt--
	indent(cnt)
	fmt.Println("==", left, right, mergeSum)
	return left + right + mergeSum
}

// merge [lo, mid] [mid+1, hi]
func merge(a []int, lo, mid, hi int) int {
	// copy array
	help := make([]int, len(a))
	for k := lo; k <= hi; k++ {
		help[k] = a[k]
	}
	// merge
	i, j := lo, mid+1
	k := lo // 犯了关键错误，这里k初始化应该是lo而不是0，不然会修改其他index的数字
	var ans int
	for i <= mid && j <= hi {
		if help[i] <= help[j] { // 产生小和
			a[k] = help[i]
			ans += help[i] * (hi - j + 1)
			i++
			k++
		} else { // help[i] > help[j]，不产生小和
			a[k] = help[j]
			j++
			k++
		}
	}
	// i > mid || j > hi
	// j<=hi || i<= mid
	for j <= hi {
		a[k] = help[j]
		j++
		k++
	}

	for i <= mid {
		a[k] = help[i]
		i++
		k++
	}
	return ans
}

// merge [lo, mid] [mid+1, hi]
func merge2(a []int, lo, mid, hi int) int {
	sum := 0
	// copy array
	help := make([]int, len(a))
	for k := lo; k <= hi; k++ {
		help[k] = a[k]
	}
	// merge
	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = help[j]
			j++
		} else if j > hi {
			a[k] = help[i]
			i++
		} else if help[i] <= help[j] { // 有小和
			a[k] = help[i]
			sum += help[i] * (hi - j + 1)
			i++
		} else {
			a[k] = help[j]
			j++
		}
	}
	return sum
}

// 自底向上的merge sort，迭代版本
func GetSmallSum3(a []int) int {
	var ans int
	n := len(a)
	for sz := 1; sz < n; sz *= 2 {
		for lo := 0; lo+sz < n; lo += 2 * sz {
			mid := lo + sz - 1
			hi := utils.MinInt(lo+2*sz-1, n-1)
			ans += merge2(a, lo, mid, hi)
		}
	}
	return ans
}

func GetSmallSum4(a []int) int {
	aux := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		aux[i] = a[i]
	}
	return smallSum2(a, aux, 0, len(a)-1)
}

func smallSum2(a, aux []int, lo, hi int) int {
	if lo >= hi {
		return 0
	}
	mid := lo + (hi-lo)/2
	// 优化1：避免复制数组
	left := smallSum2(aux, a, lo, mid)
	right := smallSum2(aux, a, mid+1, hi)
	// 优化2：检测数组书否有序
	var mergeSum int
	if aux[mid] > aux[mid+1] {
		mergeSum = merge3(a, aux, lo, mid, hi)
	} else { // 已经有序，不需要merge
		for k := lo; k <= mid; k++ {
			a[k] = aux[k]
			mergeSum += a[k] * (hi - mid)
		}
		for k := mid + 1; k <= hi; k++ {
			a[k] = aux[k]
		}
	}
	return left + right + mergeSum
}

//[lo, mid] [mid+1, hi]
func merge3(a, aux []int, lo, mid, hi int) int {
	i, j := lo, mid+1
	var ans int
	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if aux[i] <= aux[j] {
			a[k] = a[i]
			ans += aux[i] * (hi - j + 1)
			i++
		} else {
			a[k] = a[j]
			j++
		}
	}
	return ans
}

//func main() {
//	a := []int{1, 3, 5, 2, 4, 6}
//	fmt.Println(GetSmallSum(a))
//	fmt.Println(GetSmallSum2(a))
//	fmt.Println(GetSmallSum3(a))
//	fmt.Println(GetSmallSum4(a))
//}
