/*
* @Author: Yajun
* @Date:   2022/10/7 12:29
 */

package p572

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归解法：dfs枚举每个root中的节点，判断这个节点对应的子树是否和subRoot相同
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return check(root, subRoot)
	}

	return check(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

// 判断a b是否完全一样
func check(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}

	return check(a.Left, b.Left) && check(a.Right, b.Right)
}

// 将两棵树序列化，通过判断子串
// 和P101一样，单纯的前序遍历会不保留全部二叉树结构信息，需要额外引入
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

}

func preorder(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

}
