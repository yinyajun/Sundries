package main

import (
	"fmt"
	"math/rand"
)

// ------------------------------------------------------------------
// BFPRT算法
// todo: O(n)的时间内求出数组中第k小的元素
// 先用快排的partition方法代替，平均复杂度为O(N)，最坏复杂度为O(N^2)
// 首先分析下复杂度：
// partition函数：O(N)
// T = O(N) + O(N/2) +... + O(1) = O(2N) 【平均复杂度】
// T= O(N) + O(N-1) + ....~ O(N^2) 【最坏复杂度】

// 先利用普通排序方法: O(N logN)

// 利用选择排序：O(KN)

// 利用堆排序：O(NlogK) 将上一part

func selection(a []int, lo, hi int) {
	// [lo, i) sorted
	for i := lo; i <= hi; i++ {
		min := i
		for j := i + 1; j <= hi; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}

// 缺陷1：未随机化
// 缺陷2：没有对相等元素做特殊处理，大量相同元素时，性能下降
func partition1(a []int, lo, hi int) int {
	v := a[lo]
	// [lo+1, j)< v, [j, i) >= v
	j := lo + 1
	i := j
	for ; i <= hi; i++ {
		if a[i] < v {
			a[i], a[j] = a[j], a[i]
			j++
		}
	}
	// [lo+1, j)<v, [j, i)>=v
	// j-1是最后一个小于v
	a[lo], a[j-1] = a[j-1], a[lo]
	// [lo, j-1) <v, [j-1,j-1] = v,  [j, i)>=v
	return j - 1
}

// 优点：对于大量重复的元素，有可能比较均匀的分在两边
// 缺点：大量重复元素仍然会造成性能下降
func partition2(a []int, lo, hi int) int {
	//r := rand.Intn(hi-lo+1) + lo
	//a[lo], a[r] = a[r], a[lo] // 随机化pivot
	v := a[0]
	// [lo+1, i) <=v    (j, hi]>=v
	i, j := lo+1, hi
	for {
		for ; i <= hi && a[i] < v; i++ {
		} // a[i]>=v
		for ; j >= lo+1 && a[j] > v; j-- {
		} // a[j]<=v
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
	// [lo+1, i) <=v    (j, hi]>=v
	// i对应于第一个 >=v 的元素
	// j对应于第一个 <=v 的元素
	// 更主要的是，结合退出条件
	a[lo], a[j] = a[j], a[lo]
	return j
}

func partition3(a []int, lo, hi int) int {
	//r := rand.Intn(hi-lo+1) + lo
	//a[lo], a[r] = a[r], a[lo] // 随机化pivot
	v := a[0]
	// [lo+1, i] <= v   [j, hi] >= v
	i, j := lo, hi+1
	for {

		//while (a[++i] < v){
		//	if (i== hi) break;
		//}
		for i += 1; i <= hi && a[i] < v; i++ {
		} // a[i] >= v
		for j -= 1; j >= lo+1 && a[j] > v; j-- {
		} // a[j] <= v
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i] // a[i] <=v   a[j] >= v
	}
	// 注意退出条件是break， 此时 a[i] >=v , a[j] <= v
	a[lo], a[j] = a[j], a[lo]
	return j
}

// 缺陷：只有相等元素不需要交换，当数组中不存在大量重复元素时，交换次数大于二分区快排，性能下降
func threeWayPartition(a []int, lo, hi int) (int, int) {
	//r := rand.Intn(hi-lo+1) + lo
	//a[lo], a[r] = a[r], a[lo]
	v := a[lo]
	// [lo+1, lt) < v    [lt, i) == v    (gt, hi]>v
	lt, gt := lo+1, hi
	i := lo + 1
	for i <= gt {
		if a[i] < v {
			a[lt], a[i] = a[i], a[lt]
			lt++
			i++
		} else if a[i] > v {
			a[i], a[gt] = a[gt], a[i]
			gt--
		} else {
			i++
		}
	}
	a[lo], a[lt-1] = a[lt-1], a[lo]
	// [lo, lt-1) < v   [lt-1, gt+1) == v  (gt, hi] > v
	return lt - 2, gt + 1
}

// 这种写法比上面一种更好看(opt)
func threeWayPartition2(a []int, lo, hi int) (int, int) {
	//r := rand.Intn(hi-lo+1) + lo
	//a[lo], a[r] = a[r], a[lo]
	v := a[lo]
	// [lo, lt) < v    [lt, i) == v    (gt, hi]>v
	lt, gt := lo, hi
	i := lo + 1
	for i <= gt {
		if a[i] < v {
			a[lt], a[i] = a[i], a[lt]
			lt++
			i++
		} else if a[i] > v {
			a[i], a[gt] = a[gt], a[i]
			gt--
		} else {
			i++
		}
	}
	// [lo...lt)<v   [lt,gt+1) == v  (gt, hi]>v
	return lt - 1, gt + 1
}

// [lo, p]==v   (p, i] < v
// [j,q)>v   [q, hi]==v
// i, j := lo, hi+1
// p, q := lo, hi+1

// Bentley和McIlroy在dijkstra的三分区快排上改进
// 类似于上面的partition2
// 划分过程中，i遇到等于v的元素，交换到最左边
// j遇到等于v的元素，交换到最右边
// ij相遇后，再把两端与v相等的元素交换到中间
// 1. 没有重复值的时候，几乎没有额外的开销（非重复值没有额外的交换）
// 2. 有重复值的时候，虽然交换次数挺多，但这些重复值不会参与下一次排序
func BCPartition(a []int, lo, hi int) (int, int) {
	//r := median3(a, lo, hi)
	//a[r], a[lo] = a[lo], a[r]
	v := a[lo]
	// [lo, p) == v   [p, i) < v   (j, q]>v    (q, hi]==v
	p, q := lo+1, hi
	i, j := lo+1, hi

	for {
		for ; i <= hi && a[i] < v; i++ {
		} // a[i] >= v
		for ; j >= lo+1 && a[j] > v; j-- {
		} // a[j] <= v
		// 此时需要将a[i] a[j]交换到对应的位置

		// 此时有个特殊情况： i==j, 此时必然是 a[i] == a[j] == v
		// 此时的v算i的
		if i == j && a[i] == v { //
			a[i], a[p] = a[p], a[i]
			p++ // [p, i] < v
			i++ // [p, i) < v
		}
		if i > j { // pointer cross
			break
		}
		a[i], a[j] = a[j], a[i] // a[i] <=v   a[j]>=v
		if a[i] == v {
			a[i], a[p] = a[p], a[i]
			p++
		}
		if a[j] == v {
			a[j], a[q] = a[q], a[j]
			q--
		}
		// [p, i] < v    [j, q] > v
		i++
		j--
		// [p, i)<v, (j, q] > v
	}
	// 将两端的等于v的元素移动到中间, i = j + 1
	i = j + 1 // todo(有必要吗？)
	for k := lo; k < p; k++ {
		a[k], a[j] = a[j], a[k]
		j--
	}
	for k := hi; k > q; k-- {
		a[k], a[i] = a[i], a[k]
		i++
	}
	// [lo...j] < v  [i...hi]>v
	return j, i
}

func BCPartition2(a []int, lo, hi int) (int, int) {
	v := a[0]
	// [lo...p]==v    (p...i]<v    [j...q)>v     [q...hi]==v
	p, q := lo, hi+1
	i, j := lo, hi+1
	for {
		for i += 1; i <= hi && a[i] < v; i++ {
		} // a[i] >= v
		for j -= 1; j >= lo+1 && a[j] > v; j-- {
		} // a[j] <= v

		if i == j && a[i] == v { // a[i] ==  a[j] , 算i的
			p += 1
			a[i], a[p] = a[p], a[i]
		}
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
		if a[i] == v {
			p += 1
			a[i], a[p] = a[p], a[i]
		}
		if a[j] == v {
			q -= 1
			a[j], a[q] = a[q], a[j]
		}
	}
	i = j + 1 // 可能有i==j的退出情况
	for k := lo; k <= p; k++ {
		a[k], a[j] = a[j], a[k]
		j--
	}
	for k := hi; k >= q; k-- {
		a[k], a[i] = a[i], a[k]
		i++
	}
	return j, i
}

// 在数组中，采样3个元素，并取其中位数作为pivot
func median3(a []int, lo, hi int) int {
	if hi-lo+1 < 3 {
		return lo
	}
	samples := []int{
		rand.Intn(hi-lo+1) + lo,
		rand.Intn(hi-lo+1) + lo,
		rand.Intn(hi-lo+1) + lo,
	} //

	// [0, i) sorted, [i, 2] unsorted
	for i := 1; i <= 2; i++ {
		for j := i; j > 0; j-- {
			if a[samples[j]] < a[samples[j-1]] {
				samples[j], samples[j-1] = samples[j-1], samples[j]
			}
		}
	}
	return samples[1]
}

// 利用partition函数求rank=k的数
func Select(a []int, k int, partition func([]int, int, int) int) int {
	if k < 0 || k >= len(a) {
		panic("invalid arguments")
	}
	lo, hi := 0, len(a)-1
	for lo < hi {
		i := partition(a, lo, hi)
		if i < k {
			lo = i + 1
		} else if i > k {
			hi = i - 1
		} else { // i == k
			return a[i]
		}
	}
	return a[lo]
}

// three-way-partition返回
// [0...i] < k      (i..j)== k      [j...hi] > k
func Select2(a []int, k int, partition func([]int, int, int) (int, int)) int {
	if k < 0 || k >= len(a) {
		panic("invalid arguments")
	}
	lo, hi := 0, len(a)-1
	for lo < hi {
		i, j := partition(a, lo, hi)
		if k <= i {
			hi = i
		} else if k >= j {
			lo = j
		} else {
			return a[i+1]
		}
	}
	// lo == hi
	return a[lo]
}

func main() {
	a := []int{423, 54, 654, 765, 321, 654, 892, 736, 467, 12, 67, 76}
	r := partition1(a, 0, len(a)-1)
	fmt.Println(r, a[r], a)
	a = []int{423, 54, 654, 765, 321, 654, 892, 736, 467, 12, 67, 76}
	fmt.Println(Select(a, 4, partition1))

	a = []int{423, 54, 654, 765, 321, 654, 892, 736, 467, 12, 67, 76}
	r = partition2(a, 0, len(a)-1)
	fmt.Println(r, a[r], a)
	a = []int{423, 54, 654, 765, 321, 654, 892, 736, 467, 12, 67, 76}
	fmt.Println(Select(a, 4, partition2))

	a = []int{423, 54, 654, 765, 321, 654, 892, 736, 467, 12, 67, 76}
	r = partition3(a, 0, len(a)-1)
	fmt.Println(r, a[r], a)
	a = []int{423, 54, 654, 765, 321, 654, 892, 736, 467, 12, 67, 76}
	fmt.Println(Select(a, 4, partition3))

	a = []int{423, 54, 654, 765, 321, 654, 892, 736, 467, 12, 67, 76}
	r1, r2 := threeWayPartition(a, 0, len(a)-1)
	fmt.Println(r1, r2, a[r1], a[r2], a)
	a = []int{423, 54, 654, 765, 321, 654, 892, 736, 467, 12, 67, 76}
	fmt.Println(Select2(a, 4, threeWayPartition))

	a = []int{423, 54, 654, 765, 321, 654, 892, 736, 467, 12, 67, 76}
	r1, r2 = threeWayPartition2(a, 0, len(a)-1)
	fmt.Println(r1, r2, a[r1], a[r2], a)
	a = []int{423, 54, 654, 765, 321, 654, 892, 736, 467, 12, 67, 76}
	fmt.Println(Select2(a, 4, threeWayPartition2))

	a = []int{423, 54, 654, 765, 321, 654, 892, 736, 467, 12, 67, 76}
	r1, r2 = BCPartition(a, 0, len(a)-1)
	fmt.Println(r1, r2, a[r1], a[r2], a)
	a = []int{423, 54, 654, 765, 321, 654, 892, 736, 467, 12, 67, 76}
	fmt.Println(Select2(a, 4, BCPartition))

	a = []int{423, 54, 654, 765, 321, 654, 892, 736, 467, 12, 67, 76}
	r1, r2 = BCPartition2(a, 0, len(a)-1)
	fmt.Println(r1, r2, a[r1], a[r2], a)
	a = []int{423, 54, 654, 765, 321, 654, 892, 736, 467, 12, 67, 76}
	fmt.Println(Select2(a, 4, BCPartition2))

	a = []int{423, 54, 654, 765, 321, 654, 892, 736, 467, 12, 67, 76}
	selection(a, 0, len(a)-1)
	fmt.Println(a)
}
