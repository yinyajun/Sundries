package main

type node struct {
	val                 int
	left, right, parent *node
}

// 有右子树，寻找右子树中的最左节点
// 无右子树，寻找第一个比其大的父节点（不断寻找父系节点，直到找到一个父节点，它的左孩子的最右节点是该节点）
func GetNextNode(root *node) *node {
	if root == nil {
		return root
	}

	if root.right != nil {
		return leftMost(root.right)
	}
	parent := root.parent
	for parent != nil && parent.left != root {
		root = parent
		parent = root.parent
	}
	return parent // parent.left == root  || parent == nil
}

func leftMost(root *node) *node {
	cur := root
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}
