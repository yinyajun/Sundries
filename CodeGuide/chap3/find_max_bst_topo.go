package main

import (
	"CodeGuide/base/abstract"
	"CodeGuide/base/searching"
	"CodeGuide/base/utils"
	"fmt"
)

// -----------------------
// 方法1
// 时间复杂度：将所有子树都查询了一遍（O(N)）, 对于每个子树，考察了O(N)个节点，每个节点需要O(lgN)的时间
// -----------------------

// 找到二叉树中符合BST条件的最大拓扑结构
// 简化问题：以root为根的符合BST性质的最大拓扑结构

func max_bst_topo(root *abstract.TreeNode) int {
	if root == nil {
		return 0
	}
	max := _max_bst_topo(root, root)
	return utils.MaxInt(max_bst_topo(root.Left), max_bst_topo(root.Right), max)
}

// 当前节点node能为root为根的符合BST性质的最大拓扑结构所能贡献的节点数目
func _max_bst_topo(root, node *abstract.TreeNode) int {
	if root == nil || node == nil {
		return 0
	}
	if is_bst_part(root, node) {
		return _max_bst_topo(root, node.Left) +
			_max_bst_topo(root, node.Right) + 1
	}
	return 0
}

// node节点能否作为以root为根的子树的bst拓扑结构中的节点
// 换句话说，从root开始按bst的方式能否访问到node
func is_bst_part(root, node *abstract.TreeNode) bool {
	if root == nil {
		return false
	}
	if root == node {
		return true
	}
	if utils.Less(node.Key.(string), root.Key.(string)) {
		return is_bst_part(root.Left, node)
	} else {
		return is_bst_part(root.Right, node)
	}
}

// -----------------------
// 方法2
// 时间复杂度：
// -----------------------

type record struct {
	l, r int
}

func max_bst_topo2(root *abstract.TreeNode) int {
	m := make(map[*abstract.TreeNode]record)
	return _max_bst_topo2(root, m)
}

// root节点为根的拓扑贡献度，主体是个后序遍历
func _max_bst_topo2(root *abstract.TreeNode, m map[*abstract.TreeNode]record) int {
	if root == nil {
		return 0
	}

	ls := _max_bst_topo2(root.Left, m)
	rs := _max_bst_topo2(root.Right, m)

	// ls是对root左孩子的拓扑贡献度，rs是对root右孩子的拓扑贡献度
	// 现在更改为对root的拓扑贡献度，需要修改左右子树节点的record
	modify_map(root.Left, root.Key, m, true)
	modify_map(root.Right, root.Key, m, false)

	lbst, rbst := 0, 0
	if lr, ok := m[root.Left]; ok {
		lbst = lr.l + lr.r + 1
	}
	if rr, ok := m[root.Right]; ok {
		rbst = rr.l + rr.r + 1
	}
	// 将root的贡献度存储到map中
	m[root] = record{lbst, rbst}
	// root节点包含在拓扑结构中，或者在左孩子子树中，或者在右孩子子树中
	return utils.MaxInt(lbst+rbst+1, ls, rs)
}

// 当node节点对值为val的节点负责时，相应的map所需要的更改
// 返回值是修改后需要减少的贡献度，后序遍历回溯的时候，修改map中父节点对应的record
func modify_map(node *abstract.TreeNode, val interface{}, m map[*abstract.TreeNode]record, isLeft bool) int {
	if node == nil {
		return 0
	}
	if _, exist := m[node]; !exist {
		return 0
	}

	r := m[node]
	// 左子树中有节点大于val，或者右子树中有节点小于val
	// 此时将该节点断开（前序中不在探索）
	if (isLeft && utils.Less(val, node.Key)) || (!isLeft && utils.Less(node.Key, val)) {
		delete(m, node) // 需要从m中删除该node的record，相当于将对应node的record清零
		return r.l + r.r + 1
	}
	minus := 0
	// 如果位于左子树中，只要更新其右孩子节点
	if isLeft {
		minus = modify_map(node.Right, val, m, isLeft)
		r.r -= minus
	} else { // 位于右子树中，只要更新其左孩子节点
		minus = modify_map(node.Left, val, m, isLeft)
		r.l -= minus
	}
	m[node] = r
	return minus
}

func main() {
	root := searching.CreateTreeFromArray([]string{"6", "1", "0", "#", "#", "3", "#", "#", "12", "10", "4", "2",
		"#", "#", "5", "#", "#", "14", "11", "#", "#", "15", "#", "#", "13", "20", "#", "#", "16", "#", "#"})

	fmt.Println(max_bst_topo(root))
	fmt.Println(max_bst_topo2(root))
}
