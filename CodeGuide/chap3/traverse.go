package main

import (
	"CodeGuide/base/abstract"
	"CodeGuide/base/fundamentals"
	"fmt"
)

func traverse(h *abstract.TreeNode) {
	if h == nil {
		return
	}
	stack := fundamentals.NewLinkedStack()
	stack.Push(h)
	var c *abstract.TreeNode
	for !stack.IsEmpty() {
		c = stack.Peek().(*abstract.TreeNode)
		if c.Left != nil && h != c.Left && h != c.Right { // 当前节点的左右子树均未遍历过，因为可能从左右子树节点回溯
			stack.Push(c.Left)
		} else if c.Right != nil && h != c.Right {
			stack.Push(c.Right)
		} else {
			fmt.Println(stack.Pop().(*abstract.TreeNode).Key)
			h = c
		}
	}
}
