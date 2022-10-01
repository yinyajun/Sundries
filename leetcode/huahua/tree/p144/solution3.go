/*
* @Author: Yajun
* @Date:   2022/9/29 20:12
 */

package p144

// 非递归实现
// time: O(n) space:O(h)
func preorderTraversal3(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	var stack []*TreeNode

	// stack中保存的都是父节点
	for root != nil || len(stack) > 0 {
		for root != nil { // 遍历左子树直到为空
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		// 回溯到父节点，遍历右子树
		if len(stack) > 0 {
			root = stack[len(stack)-1]
			root = root.Right
			stack = stack[:len(stack)-1] // 已经不需要了
		}
	}

	return res
}
