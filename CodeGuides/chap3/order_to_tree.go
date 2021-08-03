package main

import "CodeGuide/base/abstract"

func PreIn2Tree(pre, in []int) *abstract.TreeNode {
	indexIn := make(map[int]int)
	for i, num := range in {
		indexIn[num] = i // vital
	}
	return preIn2Tree(pre, 0, len(pre)-1, in, 0, len(in)-1, indexIn)
}

func preIn2Tree(pre []int, preLo, preHi int, in []int, inLo, intHi int, indexIn map[int]int) *abstract.TreeNode {
	if preLo > preHi {
		return nil
	}
	head := abstract.NewTreeNode(pre[preLo], nil)
	idx := indexIn[pre[preLo]]
	head.Left = preIn2Tree(pre, preLo+1, preLo+idx-inLo, in, inLo, idx-1, indexIn) // preLo + 1 + (idx - inLo) -1
	head.Right = preIn2Tree(pre, preLo+idx-inLo+1, preHi, in, idx+1, intHi, indexIn)
	return head
}

func InPost2Tree(in, post []int) *abstract.TreeNode {

	indexIn := make(map[int]int)
	for i, num := range in {
		indexIn[num] = i
	}
	return inPost2Tree(in, 0, len(in)-1, post, 0, len(post)-1, indexIn)
}

func inPost2Tree(in []int, inLo, inHi int, post []int, postLo, postHi int, indexIn map[int]int) *abstract.TreeNode {
	if postLo > postHi {
		return nil
	}
	head := abstract.NewTreeNode(post[postHi], nil)
	idx := indexIn[post[postHi]]

	head.Left = inPost2Tree(in, inLo, idx-1, post, postLo, postLo+idx-inLo-1, indexIn) // post + (idx - inLo) -1
	head.Right = inPost2Tree(in, idx+1, inHi, post, postLo+idx-inLo, postHi-1, indexIn)
	return head
}

func PrePost2Tree(pre, post []int) *abstract.TreeNode {

	indexes := make(map[int]int)
	for i, num := range post {
		indexes[num] = i // post序中的index
	}
	return prePost2Tree(pre, 0, len(pre)-1, post, 0, len(post)-1, indexes)
}

// 左子树的根（前序在左段第一个）（后序在左端最后一个）
func prePost2Tree(pre []int, preLo, preHi int, post []int, postLo, postHi int, indexes map[int]int) *abstract.TreeNode {
	if preLo == preHi { // 防止后面的越界，对base case做一个修改
		return abstract.NewTreeNode(pre[preLo], nil)
	}
	head := abstract.NewTreeNode(pre[preLo], nil)
	index := indexes[pre[preLo+1]] // 特别注意，这里会越界，需要对base case做调整
	head.Left = prePost2Tree(pre, preLo+1, preLo+1+index-postLo, post, postLo, index, indexes)
	head.Right = prePost2Tree(pre, preLo+2+index-postLo, preHi, post, index+1, postHi-1, indexes)
	return head
}
