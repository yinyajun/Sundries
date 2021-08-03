package main

import "CodeGuide/base/abstract"

// 统计完全二叉树的节点个数，时间复杂度为O(h^2)

// 如果右子树达到最后一层，那么左子树一定是满二叉树（2^h-1），然后加上右子树的节点数
// 如果右子树无法达到最后一层，那么右子树一定是满二叉树，然后加上左子树的节点数

func NodeNum(root *abstract.TreeNode) int {
	if root == nil {
		return 0
	}
	return nodeNum(root, 1, mostLeftLevel(root, 1))
}

func nodeNum(root *abstract.TreeNode, lvl, h int) int {
	if lvl == h { // 叶子节点
		return 1
	}
	if mostLeftLevel(root.Right, lvl+1) == h { // Note lvl + 1
		// 左子树为满二叉树
		return 1<<(h-lvl) - 1 + 1 + nodeNum(root.Right, lvl+1, h)
	} else {
		// 右子树为满二叉树
		return nodeNum(root.Left, lvl+1, h) + 1 + 1<<(h-lvl-1) - 1
	}
}

func mostLeftLevel(root *abstract.TreeNode, lvl int) int {
	for root.Left != nil {
		root = root.Left
		lvl++
	}
	return lvl - 1
}
