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
