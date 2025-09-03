/*
* @Author: Yajun
* @Date:   2022/10/9 17:29
 */

package p437

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 穷举所有可能，访问每个节点，检测每个节点作为起始点的路径个数
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	res := rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

// root为首的和为targetSum的路径个数
func rootSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	var res int
	if targetSum == root.Val {
		res++
	}

	res += rootSum(root.Left, targetSum-root.Val)
	res += rootSum(root.Right, targetSum-root.Val)
	return res
}
