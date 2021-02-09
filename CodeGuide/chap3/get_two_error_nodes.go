package main

import (
	"CodeGuide/base/abstract"
	"CodeGuide/base/fundamentals"
	"CodeGuide/base/searching"
	"CodeGuide/base/utils"
	"fmt"
)

// 调整BST中两个错误节点

// 主要思路是，BST的中序遍历是递增序列，如果交换了两个节点，会有发生逆序
// 第一个错误节点是首次逆序的较大节点，第二个错误节点是最后一次逆序的较小节点
// 通过中序遍历的方式，找到这两个节点（需要pre，cur）
func GetTwoErrorNodes(root *abstract.TreeNode) []*abstract.TreeNode {
	// 采用非递归形式的中序遍历
	nodes := make([]*abstract.TreeNode, 2)
	stack := fundamentals.NewLinkedStack()
	var cur, pre *abstract.TreeNode
	cur = root
	for cur != nil || !stack.IsEmpty() {
		if cur != nil {
			stack.Push(cur)
			cur = cur.Left
		} else { // cur == nil
			cur = stack.Pop().(*abstract.TreeNode)
			if pre != nil && utils.Less(cur.Key, pre.Key) { // 找到逆序
				if nodes[0] == nil { // first
					nodes[0] = pre
				} else { // second
					nodes[1] = cur
				}
			}
			pre = cur
			cur = cur.Right
		}
	}
	return nodes
}

func main() {
	root := searching.CreateTreeFromArray([]string{"3", "5", "1", "#", "#", "#", "2", "4", "#", "#", "#"})
	nodes := GetTwoErrorNodes(root)
	fmt.Println(nodes[0].Key)
	fmt.Println(nodes[1].Key)
}
