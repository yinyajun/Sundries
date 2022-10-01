/*
* @Author: Yajun
* @Date:   2022/10/1 17:48
 */

package p987

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}



	fmt.Println(verticalTraversal(root))
}
