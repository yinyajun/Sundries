package main

import (
	"CodeGuide/base/abstract"
	"CodeGuide/base/fundamentals"
)

// 统计和生成所有不同的二叉树

func NumTrees(n int) int {
	if n < 2 {
		return 1
	}
	nums := make([]int, n+1)
	nums[0] = 1
	for i := 1; i < n+1; i++ {
		for j := 1; j < i+1; j++ {
			nums[i] += nums[j-1] * nums[i-j]
		}
	}
	return nums[n]
}

func generateTree(lo, hi int) *fundamentals.LinkedQueue {
	res := fundamentals.NewLinkedQueue()
	if lo > hi {
		res.Enqueue(nil)
	}
	for i := lo; i <= hi; i++ {
		head := abstract.NewTreeNode(i, nil)
		lSubs := generateTree(lo, i-1)
		rSubs := generateTree(i+1, hi)
		lSubsIter, rSubsIter := lSubs.Iterate(), rSubs.Iterate()
		for lSubsIter.First(); lSubsIter.HasNext(); {
			for rSubsIter.First(); rSubsIter.HasNext(); {
				head.Left = lSubsIter.Next().(*abstract.TreeNode)
				head.Right = rSubsIter.Next().(*abstract.TreeNode)
				res.Enqueue(cloneTree(head))
			}
		}
	}
	return res
}

func cloneTree(head *abstract.TreeNode) *abstract.TreeNode {
	if head == nil {
		return head
	}
	newHead := abstract.NewTreeNode(head.Key, head.Val)
	newHead.Left = cloneTree(head.Left)
	newHead.Right = cloneTree(head.Right)
	return newHead
}
