/*
* @Author: Yajun
* @Date:   2022/10/5 17:45
 */

package p111

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归解法：自顶向下遍历，遍历的过程中记录下最小level
func minDepth(root *TreeNode) int {
	var res = 2 << 31

	var recurse func(node *TreeNode, level int)

	recurse = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil { // leaf
			if level < res {
				res = level
			}
		}
		recurse(node.Left, level+1)
		recurse(node.Right, level+1)
	}

	recurse(root, 1)

	if res == 2<<31 {
		return 0
	}
	return res
}

// 递归解法：自底向上的，由于这里是求min，而空子树的depth=0，会影响min的计算，需要分开讨论
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// left,right is nil
	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left == nil {
		return minDepth1(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth1(root.Left) + 1
	}

	return min(minDepth1(root.Left), minDepth1(root.Right)) + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return minDepth2(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth2(root.Left) + 1
	}

	return min(minDepth2(root.Left), minDepth2(root.Right)) + 1
}
