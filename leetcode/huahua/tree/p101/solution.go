/*
* @Author: Yajun
* @Date:   2022/10/4 14:04
 */

package p101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 转化为判断中序遍历结果是否是回文
// 这个方法能通过（195/198）个case
// 遇到这种情况就无能为力了: 【1 2 2 2 # 2 #】，因为中序遍历结果并没有完全保留二叉树的结构信息
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var recurse func(node *TreeNode)
	var ser []int

	recurse = func(node *TreeNode) {
		if node == nil {
			return
		}

		recurse(node.Left)
		ser = append(ser, node.Val)
		recurse(node.Right)
	}

	recurse(root)
	// 判断ser是否是奇数长度且是回文
	if len(ser)%2 == 0 {
		return false
	}

	for i := 0; i < len(ser)/2; i++ {
		if ser[i] != ser[len(ser)-1-i] {
			return false
		}
	}
	return true
}

func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var recurse func(left, right *TreeNode) bool

	// left, right 两个子树是否对称
	// 先判断left == right，再查看 left.left和right.right是否对称，left.right和right.left是否对称
	// 也可以这么想：先判断left子树是否对应，在递归查看left.left对称否，递归查看left.right对称否
	recurse = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}

		if left == nil || right == nil {
			return false
		}

		if left.Val != right.Val {
			return false
		}

		return recurse(left.Left, right.Right) && recurse(left.Right, right.Left)
	}

	return recurse(root.Left, root.Right)
}
