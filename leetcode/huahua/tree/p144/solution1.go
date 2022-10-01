/*
* @Author: Yajun
* @Date:   2022/9/29 19:57
 */

package p144

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归实现（mutable）
func preorderTraversal(root *TreeNode) []int {
	var (
		res     []int
		recurse func(node *TreeNode)
	)

	if root == nil {
		return res
	}

	recurse = func(node *TreeNode) {
		if node == nil {
			return
		}

		res = append(res, node.Val)
		recurse(node.Left)
		recurse(node.Right)
	}

	recurse(root)
	return res
}

// 递归实现（immutable）
func preorderTraversal1(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	res = append(res, root.Val)
	res = append(res, preorderTraversal1(root.Left)...)
	res = append(res, preorderTraversal1(root.Right)...)
	return res
}
