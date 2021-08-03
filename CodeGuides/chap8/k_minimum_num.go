package main

import (
	"reflect"
)

// ------------------------------------------------------------------
// O(NlogK)的方法：维护一个大根堆，代表当前的最小的k个数
// 然后不停的遍历剩余数
// 如果数比堆顶（最小的k个数中最大的一个）大，直接忽略
// 如果数比堆顶小，用该数代替堆顶，然后使堆有序
// 【破坏了原数组】
func KMinNum1(arr []int, k int) []int {
	if k >= len(arr) {
		return arr[:]
	}
	if k <= 0 {
		return arr[:0]
	}
	// k > 0 && k <= len(arr)
	heapify1(arr, k)                // k个元素，[0, k-1]上最小的k个数
	for i := k; i < len(arr); i++ { // [0,i]上最小的k个数
		if less(arr, i, 0) { // arr[i]比堆顶小
			arr[0] = arr[i]
			sink1(arr, 0, k)
		}
	}
	return arr[:k]
}

// 根节点的索引为0, left = 2*k+1, right=2*k+2,parent= (k-1) / 2
// 如果根节点的索引为1，第一个非叶子节点为len(heap)/2
// 如果根节点的索引为0，第一个非叶子节点为(len(heap)-1-1)/2

// 堆化heap中的前k个元素
// 从首个非叶子节点开始，到根节点结束，进行堆化调整（和孩子节点比较，sink操作）
// [0, k)保持堆有序，时间复杂度为O(N)
func heapify1(heap []int, k int) {
	for i := (k - 1 - 1) / 2; i >= 0; i-- {
		sink1(heap, i, k)
	}
}

// [i, length)堆有序
// 大顶堆
func sink1(arr []int, i, length int) {
	leftChild := 2*i + 1

	for leftChild < length {
		// 父节点应该大于孩子节点
		// 如果不符合，先找到最大的孩子节点
		maxChild := leftChild
		rightChild := leftChild + 1
		if rightChild < length && less(arr, leftChild, rightChild) {
			maxChild = rightChild
		}
		if less(arr, maxChild, i) { // 父节点已经满足堆有序性质
			break
		}
		swap(arr, i, maxChild)
		i = maxChild
		leftChild = 2*i + 1
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func less(arr []int, i, j int) bool {
	if i >= len(arr) || j >= len(arr) {
		panic("out of array")
	}
	return arr[i] < arr[j]
}

// ------------------------------------------------------------------
func KMinNum2(arr []int, k int) []int {
	if k <= 0 {
		return arr[:0]
	}
	if k >= len(arr) {
		return arr[:]
	}
	// 0<k<len(arr)
	// 找k个最小，heapify [0, k-1], last non-leaf node = (k-1 -1)/2
	for i := (k - 1 - 1) / 2; i >= 0; i-- {
		sink2(arr, i, k) // [i, k-1] 堆有序
	}
	// [0, i]中最小的k个数
	for i := k; i < len(arr); i++ {
		if less(arr, i, 0) {
			arr[0] = arr[i]
			sink2(arr, 0, k)
		}
	}
	return arr[:k]
}

// [i, k) heapify
func sink2(arr []int, i, k int) {
	var child int
	for 2*i+1 < k { // left child is valid
		child = 2*i + 1
		if child+1 < k && less(arr, child, child+1) {
			child++
		}
		if !less(arr, i, child) {
			break
		}
		swap(arr, i, child)
		i = child
	}
}

// ------------------------------------------------------------------
// 破坏数组，但是元素没有丢失
type topK struct {
	length int
	k      int
	less   func(i, j int) bool
	swap   func(i, j int)
}

// return is k, [0, k) is k minimum number
// data[:k]
func NewTopK(data interface{}, k int, less func(i, j int) bool) int {
	t := &topK{
		length: reflect.ValueOf(data).Len(),
		k:      k,
		less:   less,
		swap:   reflect.Swapper(data),
	}
	return t.KMin()
}

func (t *topK) KMin() int {
	if t.k <= 0 {
		return 0
	}
	if t.length <= t.k {
		return t.length
	}
	// 0 < t.k < length
	// heapify [0, k)
	for i := (t.k - 1 - 1) / 2; i >= 0; i-- {
		t.sink(i)
	}
	// k minimum number in [0, i]
	for i := t.k; i < t.length; i++ {
		if t.less(i, 0) {
			t.swap(0, i)
			t.sink(0)
		}
	}
	return t.k
}

// [i, k)中以i为根的子树的堆化
func (t *topK) sink(i int) {
	var child int
	for 2*i+1 < t.k {
		// find max child
		child = 2*i + 1
		if child+1 < t.k && t.less(child, child+1) {
			child++
		}
		if !t.less(i, child) {
			break
		}
		t.swap(i, child)
		i = child
	}
}

//func main() {
//	arr := []int{423, 54, 654, 765, 321, 654, 892, 736, 467, 12, 67, 76}
//	k := 4
//	fmt.Println(KMinNum1(arr, k))
//	arr = []int{423, 54, 654, 765, 321, 654, 892, 736, 467, 12, 67, 76}
//	fmt.Println(KMinNum2(arr, k))
//	arr = []int{423, 54, 654, 765, 321, 654, 892, 736, 467, 12, 67, 76}
//	fmt.Println(arr[:NewTopK(arr, k, func(i, j int) bool { return arr[i] < arr[j] })])
//}
