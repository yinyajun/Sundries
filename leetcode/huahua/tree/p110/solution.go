/*
* @Author: Yajun
* @Date:   2022/10/5 13:12
 */

package p110

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归解法：类似于前序遍历，自顶向下的递归
// 时间复杂度O(n^2)，因为height函数仍然是递归的
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return abs(height(root.Left)-height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// 递归解法，自底向上，通过修改height函数，当子树不平衡的时候，直接返回-1高度
func isBalanced1(root *TreeNode) bool {
	var height func(node *TreeNode) int

	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		lHeight := height(node.Left)
		rHeight := height(node.Right)

		if abs(lHeight-rHeight) > 1 || lHeight == -1 || rHeight == -1 {
			return -1
		}
		return max(lHeight, rHeight) + 1
	}

	return height(root) >= 0
}
