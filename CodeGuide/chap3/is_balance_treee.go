package main

import (
	"CodeGuide/base/abstract"
	"CodeGuide/base/searching"
	"CodeGuide/base/utils"
	"fmt"
)

func IsBalance(root *abstract.TreeNode) bool {
	flag, h := getBalanceHeight(root, 0)
	fmt.Println(h)
	return flag
}

func getBalanceHeight(root *abstract.TreeNode, lvl int) (bool, int) {
	if root == nil {
		return true, lvl
	}
	lb, lh := getBalanceHeight(root.Left, lvl+1)
	if !lb {
		return lb, lvl
	}
	rb, rh := getBalanceHeight(root.Right, lvl+1)
	if !rb {
		return rb, lvl
	}
	height := utils.MaxInt(lh, rh)
	if (lh-rh) > 1 || (lh-rh) < -1 {
		return false, height
	}
	return true, height
}

func main() {
	root := searching.CreateTreeFromArray([]string{"3", "5", "1", "#", "#", "#", "2", "4", "#", "#", "#"})
	fmt.Println(IsBalance(root))
}
