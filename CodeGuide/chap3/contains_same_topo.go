package main

import "CodeGuide/base/abstract"

// t1是否包含t2树全部的拓扑结构

// 任何子树都可能包含t2，需要遍历t1的任何子树
func ContainsSameTopo(t1, t2 *abstract.TreeNode) bool {
	return CheckSameTopo(t1, t2) || ContainsSameTopo(t1.Left, t2) || ContainsSameTopo(t1.Right, t2)
}

// 检查t1子树是否包含t2
// 按照t2的方式前序遍历，遍历查看节点是否相同
// 时间复杂度为O(M),M为t2的节点个数
func CheckSameTopo(t1, t2 *abstract.TreeNode) bool {
	if t2 == nil {
		return true
	}
	if t1 == nil || t1.Key != t2.Key {
		return false
	}
	return CheckSameTopo(t1.Left, t2.Left) && CheckSameTopo(t1.Right, t2.Right)

}
