/*
* @Author: Yajun
* @Date:   2022/10/9 12:00
 */

package p1325

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


// 递归解法
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)
	if root.Left == nil && root.Right == nil && target == root.Val {
		return nil
	}
	return root
}

