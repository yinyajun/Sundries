package main

type node2 struct {
	val         int
	left, right *node2
}

// 找到公共祖先或者找到a元素或者找到b元素
func LCA(root, a, b *node2) *node2 {
	if root == nil {
		return nil
	}
	// 这里要注意！！！
	// case1：a，b均在以root为根的子树上，此时返回root的确是lca
	// case2：a，b中有节点不在root为根的子树上，此时返回root相当于在子树上是否能查找到a节点或者b节点
	// 根据这个特性，可以在后序的时候，完成递归
	if root == a || root == b {
		return root
	}

	left := LCA(root.left, a, b)
	right := LCA(root.right, a, b)

	if left == nil && right == nil { // 左边找不到元素 && 右边找不到元素
		return nil
	} else if right == nil { // only left!=nil
		// a b 均在左孩子
		return left
	} else if left == nil { // only right !=nil
		//a b 均在右孩子
		return right
	} else { // left != nil && right != nil
		// 左右孩子均发现 节点，说明当前节点是其lca
		return root
	}
}
