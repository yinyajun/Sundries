package main

import "CodeGuide/base/abstract"

func GenerateTree(sortedArr []int) *abstract.TreeNode {
	if len(sortedArr) == 0 {
		return nil
	}
	return _generate(sortedArr, 0, len(sortedArr)-1)
}

func _generate(arr []int, lo, hi int) *abstract.TreeNode {
	if lo > hi {
		return nil
	}
	mid := lo + (hi-lo)/2
	node := abstract.NewTreeNode(arr[mid], nil)
	node.Left = _generate(arr, lo, mid-1)
	node.Right = _generate(arr, mid+1, hi)
	return node
}
