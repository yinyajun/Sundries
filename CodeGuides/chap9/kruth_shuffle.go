package main

import (
	"math/rand"
)

//随机打印长度为N的没有重复元素的数组arr中的实现等概率随机打印arr中的M个数
// 采用knuth shuffle，保证每个位置的元素选到的概率为1/n

// 1. 从n个元素中选一个，和最后一位交换，最后1位的概率为1/n
// 2. [1..n-2]中元素n-1/n，然后随意选一个1/n-1，并将其交换到倒数第二位，倒数第二位的概率为1/n
func KnuthShuffleRandom(arr []int, m int) []int {
	if m >= len(arr) {
		m = len(arr)
	}
	for i := 0; i < m; i++ {
		r := rand.Intn(len(arr) - i)
		arr[r], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[r]
	}
	return arr[len(arr)-m:]
}

//func main() {
//	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
//	fmt.Println(KnuthShuffleRandom(arr, 3))
//}
