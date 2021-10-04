package main

import "fmt"

// N个有序数组中，从大到小打印前K个数
// 要求时间复杂度为O(KlogN)

//这道题目相当于merge多路，很容易想到堆
//而从复杂度分析的话，取k次，每次花费O(logN)的时间，也就是说需要建立一个O(N)大小的大顶堆

// left = 2i+1  right = 2i+2
// parent = (i-1)/2

type heapNode struct {
	val      int
	arrayIdx int
	valIdx   int
}

// siftdown heap[lo]
func siftdown(heap []heapNode, lo, hi int) {
	// 至少左孩子节点存在
	for 2*lo+1 <= hi {
		maxChild := 2*lo + 1
		if maxChild+1 <= hi && heap[maxChild+1].val > heap[maxChild].val { // 父节点应该大于任意一个孩子节点， 父节点需要大于最大的孩子节点， 判断右孩子节点是否是最大孩子节点
			maxChild++
		}
		if heap[lo].val >= heap[maxChild].val { // 已经符合最大堆的性质，直接退出
			break
		}
		// 不符合最大堆的性质
		heap[lo], heap[maxChild] = heap[maxChild], heap[lo]
		lo = maxChild
	}
}

// siftdown heap[lo]
// 使用赋值代替交换来优化
func siftdown2(heap []heapNode, lo, hi int) {
	e := heap[lo]
	for 2*lo+1 <= hi {
		maxChild := 2*lo + 1
		if maxChild+1 <= hi && heap[maxChild+1].val > heap[maxChild].val {
			maxChild++
		}
		if e.val >= heap[maxChild].val {
			break
		}
		heap[lo] = heap[maxChild]
		lo = maxChild
	}
	// heap[max_child] <= e || no child
	heap[lo] = e
}

// 坑点1：(hi-1)/2 >=0, 如果此时hi==0，该条件也为true；可以改为 hi-1 >= 2*0
// 坑点2：当siftup的时候，如果有个值过大，需要不断上浮，上浮的过程中会有连续交换
//	     为了优化连续交换，使用赋值代替交换。
//       在循环中，将父节点赋值到当前节点上；退出循环的时候，将最开始的值赋给当前节点

// siftup heap[hi]
func siftup(heap []heapNode, lo, hi int) {
	// 父节点存在且父节点小于当前节点
	for (hi-1) >= 2*lo && heap[(hi-1)/2].val < heap[hi].val {
		heap[(hi-1)/2], heap[hi] = heap[hi], heap[(hi-1)/2]
		hi = (hi - 1) / 2
	}
}

// siftup heap[hi]
func siftup2(heap []heapNode, lo, hi int) {
	// 至少父节点存在
	for (hi - 1) >= 2*lo {
		if heap[(hi-1)/2].val >= heap[hi].val {
			break
		}
		heap[(hi-1)/2], heap[hi] = heap[hi], heap[(hi-1)/2]
		hi = (hi - 1) / 2
	}
}

// siftup heap[hi]， 使用赋值优化连续交换
func siftup3(heap []heapNode, lo, hi int) {
	e := heap[hi]
	// 至少父节点存在
	for (hi - 1) >= 2*lo {
		if heap[(hi-1)/2].val >= e.val {
			break
		}
		heap[hi] = heap[(hi-1)/2]
		hi = (hi - 1) / 2
	}
	// heap[parent] >= e || no parent
	heap[hi] = e
}

// siftup heap[hi]， 使用赋值优化连续交换
func siftup4(heap []heapNode, lo, hi int) {
	e := heap[hi]
	for (hi-1) >= 2*lo && heap[(hi-1)/2].val < e.val {
		heap[hi] = heap[(hi-1)/2]
		hi = (hi - 1) / 2
	}
	// heap[parent] >= e || no parent
	heap[hi] = e
}

func printTopK(mat [][]int, k int) {
	if len(mat) == 0 {
		return
	}
	heapSize := len(mat)
	heap := make([]heapNode, heapSize)
	for i := 0; i < heapSize; i++ {
		heap[i] = heapNode{val: mat[i][len(mat[i])-1], arrayIdx: i, valIdx: len(mat[i]) - 1}
		siftup3(heap, 0, i)
	}

	for i := 0; i < k; i++ {
		if heapSize == 0 {
			break
		}
		fmt.Println(heap[0].val)

		if heap[0].valIdx > 0 {
			a, v := heap[0].arrayIdx, heap[0].valIdx
			heap[0].val = mat[a][v-1]
			heap[0].valIdx--
		} else {
			heap[0], heap[heapSize-1] = heap[heapSize-1], heap[0]
			heapSize--
		}
		siftdown2(heap, 0, heapSize-1)
	}
}

//func main() {
//	//arr := []int{3, 4, 2, 8, 5, 7, 1}
//	//heap := make([]heapNode, len(arr))
//	//for i := 0; i < len(arr); i++ {
//	//	heap[i] = heapNode{val: arr[i], arrayIdx: 0, valIdx: 0}
//	//	siftup4(heap, 0, i)
//	//}
//	//
//	//heap[0].val = 0
//	//siftdown2(heap, 0, len(heap)-1)
//	//
//	//for i := 0; i < len(heap); i++ {
//	//	fmt.Println(heap[i].val)
//	//}
//	mat := [][]int{
//		{2, 3, 9},
//		{4, 6, 10},
//		{1, 5, 7},
//	}
//	printTopK(mat, 100)
//
//}
