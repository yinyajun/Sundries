package main

import (
	"CodeGuide/base/abstract"
	"CodeGuide/base/searching"
	"CodeGuide/base/utils"
	"fmt"
)

func MaxDistance(root *abstract.TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	lMaxDist, lRootDist := MaxDistance(root.Left)
	rMaxDist, rRootDist := MaxDistance(root.Right)

	maxDist := utils.MaxInt(lMaxDist, rMaxDist, lRootDist+rRootDist+1)
	rootDist := utils.MaxInt(rRootDist, lRootDist) + 1
	return maxDist, rootDist
}

func main() {
	root := searching.CreateTreeFromArray([]string{
		"1", "2", "4", "#", "#", "5", "#", "#", "3", "6", "#", "#", "7", "#", "#",
	})
	fmt.Println(MaxDistance(root))
}
