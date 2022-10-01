/*
* @Author: Yajun
* @Date:   2022/9/30 10:48
 */

package p145

// 非递归实现
func postorderTraversal2(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	var stack []*TreeNode
	var last *TreeNode // 最近一次访问过的子树

	// stack中保存的都是父节点
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else { // root == nil
			root = stack[len(stack)-1] // 此时不急着出栈，需要右子树都被访问过才需要出栈
			// 难点是如何判断右子树被访问过
			if root.Right == nil || root.Right == last { // 无右子树或者右子树已经访问过
				res = append(res, root.Val)
				stack = stack[:len(stack)-1] // root节点可以出栈了
				last = root                  // 缓存最近一次访问过的子树
				root = nil                   // 无法索引下一个节点了，需要从stack中pop出下一个节点
			} else {
				root = root.Right
			}
		}
	}

	return res
}

// 非递归实现
func postorderTraversal3(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	var stack []*TreeNode
	var last *TreeNode // 最近一次访问过的子树

	// stack中保存的都是父节点
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) > 0 {
			root = stack[len(stack)-1] // 此时不急着出栈，需要右子树都被访问过才需要出栈
			// 难点是如何判断右子树被访问过
			if root.Right == nil || root.Right == last { // 无右子树或者右子树已经访问过
				res = append(res, root.Val)
				last = root // 缓存最近一次访问过的子树
				root = nil  // 注意，不然会导致一直循环
				stack = stack[:len(stack)-1]
			} else {
				root = root.Right
			}
		}
	}

	return res
}

func postorderTraversal4(root *TreeNode) (res []int) {
	var stack []*TreeNode
	var prev *TreeNode

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) > 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1] // 直接出栈

			if root.Right == nil || root.Right == prev {
				res = append(res, root.Val)
				prev, root = root, nil
			} else { // 如果需要遍历右子树，那么还需要将root塞回stack中，等右子树遍历完才能访问
				stack = append(stack, root)
				root = root.Right
			}
		}
	}
	return res
}
