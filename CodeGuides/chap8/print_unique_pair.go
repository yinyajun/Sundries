package main

import (
	"fmt"
	"sort"
)

// 这里的方法：会重复打印相同的二元组
func PrintUniquePair(a []int, k int) {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	for i, j := 0, len(a)-1; i < j; {
		s := a[i] + a[j]
		if s > k {
			j--
		} else if s < k {
			i++
		} else {
			fmt.Println(a[i], a[j])
			i++
			j--
		}
	}
}

// 需要添加一个检查 a[left]==a[left-1]? 这样重复二元组就可以省略了
func PrintUniquePair2(a []int, k int) {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	var i, j int
	i, j = 0, len(a)-1
	for i < j {
		s := a[i] + a[j]
		if s > k {
			j--
		} else if s < k {
			i++
		} else {
			if i == 0 || a[i] != a[i-1] {
				fmt.Println(a[i], a[j])
			}
			i++
			j--
		}
	}
}

func printUniqueTriad(a []int, k int) {
	if len(a) < 3 {
		return
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	for i := 0; i < len(a)-2; i++ {
		if i == 0 || a[i] != a[i-1] { // 保证第一个元素不重复
			printRest(a, k-a[i], i, i+1, len(a)-1)
		}
	}
}

func printRest(a []int, k, first, left, right int) {
	for left < right {
		s := a[left] + a[right]
		if s < k {
			left++
		} else if s > k {
			right--
		} else {
			if left == first+1 || a[left] != a[left-1] {
				fmt.Println(a[first], a[left], a[right])
			}
			left++
			right--
		}
	}
}

//func main() {
//	a := []int{-8, -4, -3, 0, 1, 1, 2, 4, 5, 8, 9, 9}
//	PrintUniquePair(a, 10)
//	PrintUniquePair2(a, 10)
//	printUniqueTriad(a, 10)
//}
