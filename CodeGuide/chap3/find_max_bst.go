package main

import (
	"CodeGuide/base/abstract"
	"CodeGuide/base/searching"
	"CodeGuide/base/utils"
	"fmt"
	"strconv"
)

func MaxBST(root *abstract.TreeNode) *abstract.TreeNode {
	min, max, size, node := postBST(root)
	fmt.Println(min, max, size)
	return node
}

const MAXVALUE = 1<<31 - 1
const MINVALUE = -1 << 31

// 思路
// 先判断子树是否是符合BST：左右子树都是BST，且root大于左子树所有值，root小于右子树所有值
// 需要等待子树的结果，然后才能得到根的结果，很明显，这是一个后序遍历的过程

func postBST(root *abstract.TreeNode) (min, max, size int, node *abstract.TreeNode) {
	if root == nil {
		return MAXVALUE, MINVALUE, 0, nil
	}
	lMin, lMax, lSize, lBST := postBST(root.Left)
	rMin, rMax, rSize, rBST := postBST(root.Right)

	rootVal, _ := strconv.Atoi(root.Key.(string))
	// 为什么这里只要lmin和rootval比较？主要是为了root节点能构成BST的情况准备的
	// 如果root节点不能构成BST，那么往上也不能构成BST，无所谓最小值和最大值了
	// 由于root节点可以构成BST的情况下，min=左子树最小, max=右子树最大
	// 考虑从base情况回溯，min可能非常大，max可能非常小，所以min=更小(lMin, rootValue), max同理
	min = utils.MinInt(lMin, rootVal)
	max = utils.MaxInt(rMax, rootVal)

	// 当前节点是否符合BST
	if root.Left == lBST && root.Right == rBST && utils.Less(rootVal, rMin) && utils.Less(lMax, rootVal) {
		return min, max, lSize + rSize + 1, root
	}
	// 当前节点不符合BST，那么BST只可能是lBST和rBST中最大的那个
	if lSize > rSize {
		return min, max, lSize, lBST
	} else {
		return min, max, rSize, rBST
	}
}

func main() {
	root := searching.CreateTreeFromArray([]string{"6", "1", "0", "#", "#", "3", "#", "#", "12", "10", "4", "2",
		"#", "#", "5", "#", "#", "14", "11", "#", "#", "15", "#", "#", "13", "20", "#", "#", "16", "#", "#"})
	node := MaxBST(root)
	fmt.Println(node)
}
