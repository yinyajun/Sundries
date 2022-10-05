/*
* @Author: Yajun
* @Date:   2022/10/5 11:31
 */

package p101

func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var stack []*TreeNode

	stack = append(stack, root.Left, root.Right)

	for len(stack) > 0 {
		right, left := stack[len(stack)-1], stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		if left == nil && right == nil {
			continue
		}

		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		stack = append(stack, left.Left, right.Right)
		stack = append(stack, left.Right, right.Left)
	}
	return true
}
