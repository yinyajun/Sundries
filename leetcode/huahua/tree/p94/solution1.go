/*
* @Author: Yajun
* @Date:   2022/9/29 17:59
 */

package p94

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归实现（mutable）
func inorderTraversal(root *TreeNode) []int {
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

		recurse(node.Left)
		res = append(res, node.Val)
		recurse(node.Right)
	}

	recurse(root)
	return res
}

// 递归实现
func inorderTraversal1(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	res = append(res, inorderTraversal1(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal1(root.Right)...)
	return res
}
