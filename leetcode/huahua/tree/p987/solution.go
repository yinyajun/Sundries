/*
* @Author: Yajun
* @Date:   2022/10/1 17:14
 */

package p987

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalTraversal(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var aux = make(map[int][][]int)
	var minCol, maxCol int
	var recurse func(node *TreeNode, r, c int)

	recurse = func(node *TreeNode, r, c int) {
		if node == nil {
			return
		}

		if c < minCol {
			minCol = c
		}
		if c > maxCol {
			maxCol = c
		}

		_, ok := aux[c]
		if !ok {
			aux[c] = [][]int{}
		}

		fmt.Println(c,r, aux[c])

		if r >= len(aux[c]) {
			aux[c] = append(aux[c], []int{})
		}
		fmt.Println(c,r, aux[c])

		aux[c][r] = append(aux[c][r], root.Val)

		recurse(node.Left, r+1, c-1)
		recurse(node.Left, r+1, c+1)
	}

	for i := minCol; i <= maxCol; i++ {
		if len(aux[i]) == 0 {
			continue
		}
		var temp []int
		for j := 0; j < len(aux[i]); j++ {
			sort.Slice(aux[i][j], func(a, b int) bool {
				return aux[i][j][a] < aux[i][j][b]
			})
			temp = append(temp, aux[i][j]...)
		}
		res = append(res, temp)
	}

	recurse(root, 0, 0)
	return res
}
