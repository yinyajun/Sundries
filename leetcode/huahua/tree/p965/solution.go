/*
* @Author: Yajun
* @Date:   2022/10/8 18:31
 */

package p965

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil {
		return root.Val == root.Right.Val && isUnivalTree(root.Right)
	}
	if root.Right == nil {
		return root.Val == root.Left.Val && isUnivalTree(root.Left)
	}

	return root.Val == root.Left.Val && root.Val == root.Right.Val &&
		isUnivalTree(root.Left) && isUnivalTree(root.Right)
}

func isUnivalTree1(root *TreeNode) bool {
	return root == nil ||
		(root.Left == nil || root.Val == root.Left.Val && isUnivalTree1(root.Left)) &&
			(root.Right == nil || root.Val == root.Right.Val && isUnivalTree1(root.Right))
}


