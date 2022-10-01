/*
* @Author: Yajun
* @Date:   2022/9/30 15:42
 */

package p145

// 前序：【中，左，右】
// 后序：【左，右，中】
// 逆后序：【中，右，左】
// 逆后序和前序，只是遍历子树的顺序反了，完全可以转为类似前序的方式处理（而前序的迭代解法很简单）
func postorderTraversal5(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var stack []*TreeNode
	stack = append(stack, root)

	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, root.Val)

		if root.Left != nil {
			stack = append(stack, root.Left)
		}

		if root.Right != nil {
			stack = append(stack, root.Right)
		}
	}

	// reverse
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}
