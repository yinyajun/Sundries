package main

import (
	"CodeGuide/base/abstract"
	"CodeGuide/base/fundamentals"
	"CodeGuide/base/searching"
	"fmt"
)

func levelPrint(root *abstract.TreeNode) {
	if root == nil {
		return
	}
	queue := fundamentals.NewLinkedQueue()
	queue.Enqueue(root)
	lvl := 1
	for !queue.IsEmpty() {
		size := queue.Size()
		fmt.Print("Level ", lvl, " : ")
		lvl++
		for i := 0; i < size; i++ {
			node := queue.Dequeue().(*abstract.TreeNode)
			fmt.Print(node.Key, " ")
			if node.Left != nil {
				queue.Enqueue(node.Left)
			}
			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}
		fmt.Println()
	}
}

func zigzagPrint2(root *abstract.TreeNode) {
	if root == nil {
		return
	}
	queue := fundamentals.NewDeque()
	lvl := 1
	queue.PushBack(root)
	for !queue.IsEmpty() {
		size := queue.Size()
		printLevelAndDirection(lvl)
		for i := 0; i < size; i++ {
			if lvl%2 == 1 {
				node := queue.PopFront().(*abstract.TreeNode)
				fmt.Print(node.Key, " ")
				if node.Left != nil {
					queue.PushBack(node.Left)
				}
				if node.Right != nil {
					queue.PushBack(node.Right)
				}
			} else {
				node := queue.PopBack().(*abstract.TreeNode)
				fmt.Print(node.Key, " ")
				if node.Right != nil {
					queue.PushFront(node.Right)
				}
				if node.Left != nil {
					queue.PushFront(node.Left)
				}
			}
		}
		lvl++
		fmt.Println()
	}
}

func printLevelAndDirection(lvl int) {
	fmt.Print("Level ", lvl)
	if lvl%2 == 1 {
		fmt.Print(" from left to right: ")
	} else {
		fmt.Print(" from right to left: ")
	}
}

func main() {
	root := searching.CreateTreeFromArray([]string{"1", "2", "4", "#", "#", "#", "3", "5", "7", "#", "#", "8", "#", "#", "6", "#", "#"})
	levelPrint(root)
	fmt.Println("-------------------------")
	zigzagPrint2(root)
}
