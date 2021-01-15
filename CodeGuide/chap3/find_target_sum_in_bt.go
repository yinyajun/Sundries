package main

import (
	"CodeGuide/base/abstract"
	"CodeGuide/base/searching"
	"CodeGuide/base/utils"
	"strconv"
)

// 在二叉树中找到累加和为指定值的最长路径长度
// 时间复杂度O(N)，空间复杂度O(h)，N为节点数目，h为树的高度

// 注意边界条件，m[0]=0, 那么level从1开始

func MaxLengthTree(root *abstract.TreeNode, target int) int {
	m := make(map[int]int)
	m[0] = 0
	maxLen, preSum, level := 0, 0, 1
	PreOrderSum(root, target, preSum, level, m, &maxLen)
	return maxLen
}

func value(v interface{}) int {
	num, _ := strconv.Atoi(v.(string))
	return num
}

func PreOrderSum(root *abstract.TreeNode, target, preSum, level int, m map[int]int, maxLen *int) {
	if root == nil {
		return
	}
	// root != nil
	curSum := preSum + value(root.Key)
	if preLevel, ok := m[curSum-target]; ok {
		*maxLen = utils.MaxInt(*maxLen, level-preLevel)
	}
	if _, ok := m[curSum]; !ok {
		m[curSum] = level
	}
	PreOrderSum(root.Left, target, curSum, level+1, m, maxLen)
	PreOrderSum(root.Right, target, curSum, level+1, m, maxLen)

	// 回溯
	if m[curSum] == level {
		delete(m, curSum)
	}
}

func main() {
	root := searching.CreateTreeFromArray([]string{
		"1", "2", "#", "4", "7", "#", "#", "8", "#", "11", "13", "#", "#", "14", "#", "#",
		"3", "5", "9", "12", "15", "#", "#", "16", "#", "#", "#", "10", "#", "#", "6", "#", "#"})
	println(MaxLengthTree(root, 9))
}
