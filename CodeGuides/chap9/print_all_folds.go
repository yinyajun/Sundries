package main

import "fmt"

// 打印折纸纸印
// N=1             down
// N=2      up            down
// N=3   up    down     up      down

// 规律如下：
// 第i+1次产生的折痕，是在第i次的基础上，分别插入上下折痕
// 满二叉树，左孩子为上折痕，右孩子为下折痕
// 为了打印所有折痕，就是中序遍历的过程

func printAllFolds(n int) {
	printFolds(1, n, true)
}

// 空间复杂度为O(N), 树的高度
func printFolds(i, n int, down bool) {
	if i > n {
		return
	}
	printFolds(i+1, n, false)
	if down {
		fmt.Print("down ")
	} else {
		fmt.Print("up ")
	}
	printFolds(i+1, n, true)
}

//func main() {
//	printAllFolds(3)
//}
