/*
* @Author: Yajun
* @Date:   2022/10/7 12:29
 */

package p572

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归解法：dfs枚举每个root中的节点，判断这个节点对应的子树是否和subRoot相同
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return check(root, subRoot)
	}

	return check(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

// 判断a b是否完全一样
func check(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}

	return check(a.Left, b.Left) && check(a.Right, b.Right)
}

// 将两棵树序列化，通过判断子串
// 和P101一样，单纯的前序遍历会不保留全部二叉树结构信息，需要额外引入信息，保证遍历结果的唯一性
func isSubtree2(root *TreeNode, subRoot *TreeNode) bool {
	// 判断sl是否是rl的子序列：kmp
	return kmp(preorder(root), preorder(subRoot))
}

func preorder(root *TreeNode) []int {
	var res []int
	var recurse func(node *TreeNode)
	const (
		lNil = 2 << 31
		rNil = 2<<31 + 1
	)

	recurse = func(node *TreeNode) {
		if node == nil {
			return
		}

		res = append(res, node.Val)
		if node.Left == nil {
			res = append(res, lNil)
		}
		if node.Right == nil {
			res = append(res, rNil)
		}

		recurse(node.Left)
		recurse(node.Right)
	}

	recurse(root)
	return res
}

func kmp(text, pattern []int) bool {
	// 求解next数组
	// next[i] = pre   -->  pattern[0...pre-1] = pattern[i-pre...i-1]
	// if pattern[pre] == pattern[i]
	// 		next[i+1] = pre+1 = next[i] + 1
	// else
	// 		for pattern[i] != pattern[pre] {pre = next[pre] }
	//      next[i+1] = next[pre] + 1
	getNext := func(pattern []int) []int {
		var pre int
		var next = make([]int, len(pattern)+1)

		// 该定义下，next数组前2位必然是0
		for i := 1; i < len(pattern); i++ {
			for pattern[i] != pattern[pre] && pre != 0 { // 不停地寻找可以和pattern[i]组成的最大公共前后缀
				pre = next[pre]
			}
			// pre == 0 || pattern[i] == pattern[pre]
			if pattern[pre] == pattern[i] {
				pre += 1
			}
			next[i+1] = pre
		}
		return next
	}

	next := getNext(pattern)
	var j int // pattern 指针
	for i := 0; i < len(text); i++ {
		for pattern[j] != text[i] && j > 0 { // 查询next数组，只改变pattern的起点
			j = next[j]
		}

		if pattern[j] == text[i] {
			j++
		}
		if j == len(pattern) {
			return true
		}
	}
	return false
}
