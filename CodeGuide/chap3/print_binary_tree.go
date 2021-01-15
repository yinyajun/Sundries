package main

import (
	"CodeGuide/base/abstract"
	"CodeGuide/base/searching"
	"fmt"
	"strings"
)

func PrintBinaryTree(root *abstract.TreeNode) {
	printInOrder(root, 0, "H")
}

func printInOrder(root *abstract.TreeNode, height int, sign string) {
	if root == nil {
		return
	}
	printInOrder(root.Right, height+1, "v")

	val := sign + root.Key.(string) + sign
	lenL := (17 - len(val)) / 2
	lenR := 17 - len(val) - lenL
	val = getSpace(lenL) + val + getSpace(lenR)
	fmt.Println(getSpace(height*17) + val)

	printInOrder(root.Left, height+1, "^")

}

func getSpace(num int) string {
	buf := strings.Builder{}
	for i := 0; i < num; i++ {
		buf.WriteString(" ")
	}
	return buf.String()
}

func main() {
	root := searching.CreateTreeFromArray([]string{"6", "1", "0", "#", "#", "3", "#", "#", "12", "10", "4", "2",
		"#", "#", "5", "#", "#", "14", "11", "#", "#", "15", "#", "#", "13", "20", "#", "#", "16", "#", "#"})
	PrintBinaryTree(root)
}
