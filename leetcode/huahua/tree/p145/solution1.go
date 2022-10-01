/*
* @Author: Yajun
* @Date:   2022/9/30 10:44
 */

package p145

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归实现(mutable)
func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var recurse func(node *TreeNode)

	recurse = func(node *TreeNode) {
		if node == nil {
			return
		}
		recurse(node.Left)
		recurse(node.Right)
		res = append(res, node.Val)
	}

	recurse(root)
	return res
}

// author: Huahua
// 递归（immutable）
// post([4 1 # 3 2]) = post([1]) + post([3 2]) + [4]
// post([1]) = [1]
// post([3 2]) = post([2]) + [3] = [2 3]
// post([4 1 # 3 2]) = [1 2 3 4]
func postorderTraversal1(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	res = append(res, postorderTraversal1(root.Left)...)
	res = append(res, postorderTraversal1(root.Right)...)
	res = append(res, root.Val)
	return res
}
