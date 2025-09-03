/*
* @Author: Yajun
* @Date:   2022/10/9 00:18
 */

package p669

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归解法
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > high {
		return trimBST(root.Left, low, high)
	}

	if root.Val < low {
		return trimBST(root.Right, low, high)
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

// 迭代解法（todo: repeat）
func trimBST1(root *TreeNode, low, high int) *TreeNode {
	// 首先确定在区间内的根节点
	for root != nil && (root.Val < low || root.Val > high) {
		if root.Val < low {
			root = root.Right
		} else {
			root = root.Left
		}
	}

	// root ==nil || low<=root.val<=high
	if root == nil {
		return nil
	}

	// trim root.left
	for node := root; node.Left != nil; {
		if node.Left.Val < low { // 如果它的左结点left的值小于low，那么left以及left的左子树都不符合要求，我们将node的左结点设为left的右结点，然后再重新对node的左子树进行修剪。
			node.Left = node.Left.Right
		} else { // 如果它的左结点left的值大于等于low，又因为node的值已经符合要求，所以left的右子树一定符合要求。基于此，我们只需要对left的左子树进行修剪。我们令node等于left ，然后再重新对node 的左子树进行修剪。
			node = node.Left
		}

	}

	// trim root.right
	for node := root; node.Right != nil; {
		if node.Right.Val > high {
			node.Right = node.Right.Left
		} else {
			node = node.Right
		}
	}
	return root
}
