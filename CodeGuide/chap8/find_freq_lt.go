/*
* 在数组中，打印其中出现次数大于一半的数 【O(N)时间，O(1)空间】
* 在数组中，再给一个整数K，打印所有出现次数大于N/K的数 【O(NK)时间，O(K)空间】
 */

package main

import "fmt"

// 每次删除两个不同的数，如果一个数出现次数大于一半
// 最后留下的数肯定是该数（当然也需要检验一下）
func printHalfMajor(a []int) {
	if len(a) <= 0 {
		fmt.Println("arr length too small")
	}
	cand := 0
	times := 0

	for i := 0; i < len(a); i++ {
		if times == 0 { // 没有候选者
			cand = a[i]
			times++
		} else if cand == a[i] { // 有候选者，但是当前的数和候选者一样
			times++
		} else { // 有候选者且当前的数和候选者不一样
			times--
		}
	}
	// 当遍历过去，且times--，此时才算是删除一对
	// 而times++，代表了当前无法构成一对，先欠着（数放在cand中）

	// 验证结果
	times = 0
	for i := 0; i < len(a); i++ {
		if a[i] == cand {
			times++
		}
	}
	if times > len(a)/2 {
		fmt.Println(cand)
		return
	}
	fmt.Println("no such number.")
}

func printKMajor(a []int, k int) {
	if len(a) <= 0 || k <= 1 {
		fmt.Println("invalid arguments")
	}
	cand := make(map[int]int)

	for i := 0; i < len(a); i++ {
		if _, ok := cand[a[i]]; ok {
			cand[a[i]]++
		} else if len(cand) < k-1 {
			cand[a[i]]++
		} else { // 已经有k-1个候选者
			minusOnce(cand)
		}
	}
	reals := make(map[int]int)
	for i := 0; i < len(a); i++ {
		if _, ok := cand[a[i]]; ok {
			reals[a[i]]++
		}
	}

	hasPrint := false
	for i, j := range reals {
		if j > len(a)/k {
			fmt.Print(i)
			hasPrint = true
			fmt.Print(" ")
		}
	}
	if !hasPrint {
		fmt.Println("no such number.")
	}
	fmt.Println()
}

func minusOnce(cand map[int]int) {
	// 尽量不要在循环中删除元素
	toDelete := []int{}
	for k, v := range cand {
		if v == 1 {
			toDelete = append(toDelete, k)
		} else {
			cand[k]--
		}
	}
	for _, k := range toDelete {
		delete(cand, k)
	}
}

//func main() {
//	a := []int{4, 2, 3, 4, 2, 3, 4, 4, 4}
//	printHalfMajor(a)
//
//	b := []int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 1, 3}
//	printKMajor(b, 3)
//}
