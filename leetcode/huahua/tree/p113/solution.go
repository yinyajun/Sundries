/*
* @Author: Yajun
* @Date:   2022/10/9 14:51
 */

package p113

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归解法：类似回溯
func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var sum int
	var path []int

	var recurse func(node *TreeNode)

	recurse = func(node *TreeNode) {
		if node == nil {
			return
		}

		sum += node.Val
		path = append(path, node.Val)

		if node.Left == nil && node.Right == nil && sum == targetSum {
			var pathCopied = make([]int, len(path))
			for i, n := range path {
				pathCopied[i] = n
			}
			res = append(res, pathCopied)
			return
		}

		if node.Left != nil {
			recurse(node.Left)
			sum -= node.Left.Val
			path = path[:len(path)-1]
		}

		if node.Right != nil {
			recurse(node.Right)
			sum -= node.Right.Val
			path = path[:len(path)-1]
		}
	}

	recurse(root)
	return res
}

// 递归解法：类似回溯
func pathSum1(root *TreeNode, targetSum int) [][]int {
	var res [][]int

	if root == nil {
		return res
	}

	var recurse func(node *TreeNode, sum int, path []int)

	recurse = func(node *TreeNode, sum int, path []int) {
		if node.Left == nil && node.Right == nil && targetSum == sum {
			var pathCopied = make([]int, len(path))
			for i, n := range path {
				pathCopied[i] = n
			}
			res = append(res, pathCopied)
			return
		}

		if node.Left != nil {
			recurse(node.Left, sum+node.Left.Val, append(path, node.Left.Val))
		}

		if node.Right != nil {
			recurse(node.Right, sum+node.Right.Val, append(path, node.Right.Val))
		}
	}

	recurse(root, root.Val, []int{root.Val})
	return res
}

// 递归解法：类似回溯
func pathSum2(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var sum int
	var path []int
	var recurse func(node *TreeNode)

	recurse = func(node *TreeNode) {
		if node == nil {
			return
		}

		sum += node.Val
		path = append(path, node.Val)
		defer func() {
			sum -= node.Val
			path = path[:len(path)-1]
		}()

		if node.Left == nil && node.Right == nil && targetSum == sum {
			res = append(res, append([]int{}, path...))
		}

		recurse(node.Left)
		recurse(node.Right)
	}

	recurse(root)
	return res
}

type Pair struct {
	node *TreeNode
	sum  int
}

// BFS解法 (todo repeat)
func pathSum3(root *TreeNode, targetSum int) [][]int {
	var res [][]int

	if root == nil {
		return res
	}

	parent := map[*TreeNode]*TreeNode{}
	getPath := func(node *TreeNode) []int {
		var path []int
		for node != nil {
			path = append(path, node.Val)
			node = parent[node]
		}

		for i := 0; i < len(path)/2; i++ {
			path[i], path[len(path)-1-i] = path[len(path)-1-i], path[i]
		}
		return path
	}

	var queue []Pair
	queue = append(queue, Pair{root, 0})

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		sum := p.sum + p.node.Val
		if p.node.Left == nil && p.node.Right == nil && sum == targetSum {
			res = append(res, getPath(p.node))
			continue
		}

		if p.node.Left != nil {
			parent[p.node.Left] = p.node
			queue = append(queue, Pair{node: p.node.Left, sum: sum})
		}
		if p.node.Right != nil {
			parent[p.node.Right] = p.node
			queue = append(queue, Pair{node: p.node.Right, sum: sum})
		}
	}
	return res
}
