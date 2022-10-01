/*
* @Author: Yajun
* @Date:   2022/9/29 18:10
 */

package p94

// 非递归实现
// 沿着左子树不断访问，直到为空，回溯到之前节点，然后转向其右子树，然后重复。
// time: O(n) space:O(h)
func inorderTraversal2(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	var stack []*TreeNode

	// 栈中保存的都是父节点，用于回溯时访问并索引到其右子树
	// 当栈中待回溯的节点，或者当前节点不为空，都能继续去遍历
	for root != nil || len(stack) > 0 {
		if root != nil { // 沿着左子树方向入栈
			stack = append(stack, root)
			root = root.Left
		} else { // 左子树为空，回溯到父节点，访问之，然后找到其右节点（然后就没有利用价值了，可以出栈了）
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, root.Val)
			root = root.Right
		}
	}

	return res
}

// time: O(n) space:O(h)
func inorderTraversal3(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	var stack []*TreeNode

	for root != nil || len(stack) > 0 {
		// 优化为连续循环
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 此时root==nil，循环发生了多次，但是循环变量只保证单次循环满足条件
		if len(stack) > 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, root.Val)
			root = root.Right
		}
	}

	return res
}
