package main

import "fmt"

// 时间复杂度要求O(N)
// 空间复杂度要求O(1)

// 空间复杂度为O(N)
func StringOnlyOnce(a string) bool {
	record := make(map[byte]struct{})
	for i := range a {
		if _, ok := record[a[i]]; ok { // see you again
			return false
		} else {
			record[a[i]] = struct{}{} // the first glance
		}
	}
	return true
}

func StringOnlyOnce2(a string) bool {
	record := make([]bool, 256)
	for i := range a {
		if record[a[i]] == true {
			return false
		} else {
			record[a[i]] = true
		}
	}
	return true
}

// 空间复杂度要求O(1)，要求先排序
// 最后考察的就是使用O(1)的排序方法
// O(N)的排序方法依赖额外空间
// O(NLogN)中merge，quick的空间复杂度都不止O(1)
// 使用heapsort

func StringOnlyOnce3(a string) bool {
	c := []byte(a)
	HeapSort(c)
	for i := 0; i < len(c); i++ {
		if i > 0 && c[i] == c[i-1] {
			return false
		}
	}
	return true
}

func HeapSort(a []byte) {
	Heapify(a)
	for i := len(a) - 1; i >= 1; i-- {
		a[0], a[i] = a[i], a[0] // 将最大元素放到最后面
		sink(a, 0, i-1)
	}
}

func Heapify(a []byte) {
	n := len(a)
	for i := (n - 1 - 1) / 2; i >= 0; i-- {
		sink(a, i, n-1)
	}
}

// 大顶堆，将小的元素下沉下去
// left = 2*i+1, right = 2*i+2 , parent = (i-1)/2
func sink(a []byte, lo, hi int) {
	root := lo
	for 2*root+1 <= hi { // 至少有叶子
		j := 2*root + 1
		// 找出左右子树中最大的那个
		if j+1 <= hi && a[j+1] > a[j] {
			j++
		}
		// 如果满足堆性质(根大)，则拉倒
		if a[root] >= a[j] {
			break
		}
		// 将root和最大的孩子交换
		a[root], a[j] = a[j], a[root]
		root = j
	}
}

func main() {
	fmt.Println(StringOnlyOnce("121"))
	fmt.Println(StringOnlyOnce2("121"))
	fmt.Println(StringOnlyOnce3("121"))
}
