package main

// 用树来实现二叉堆

type _node struct {
	val         int
	left, right *_node
	parent      *_node
}

type MyHeap struct {
	head *_node
	last *_node
	size int
	less func(i, j int) bool
}

func (h *MyHeap) getSize() int { return h.size }

func (h *MyHeap) getHead() int {
	if h.head != nil {
		return h.head.val
	}
	panic("head is nil")
}

// 1. size=0，直接添加
// 2. size>0，有堆节点
// 	1. last是当前层最后一个节点，在新层最左位置添加
// 	2. last是左孩子，在右孩子处添加
//  3. last是右孩子，比较复杂，要向上寻找，寻找下一个节点的父节点
//          想上寻找last的祖先，该祖先节点ans为左孩子，那么ans的右边兄弟子树的最左边节点的左孩子就是要插入的位置
//          证明：ans为左孩子，子树的最右边节点为last，那么ans的右兄弟的子树的最左边节点就是空的，其左孩子节点可以插入

func (h *MyHeap) add(val int) {
	node := &_node{val: val}

	if h.size == 0 {
		h.head = node
		h.last = node
		h.size++
		return
	}

	cur := h.last
	par := cur.parent

	for par != nil && par.left != cur { // 只要cur不是左孩子就一直向上，直到寻找到某个祖先节点为左孩子或者找到根节点
		cur = par
		par = cur.parent
	}
	// par == nil || par.left == cur

	var nodeToAdd *_node // 将要添加子节点的节点

	if par == nil { // 已经找到根节点，说明last为当前层的最后一个节点
		nodeToAdd = mostLeft(cur)
		nodeToAdd.left = node
		node.parent = nodeToAdd
	} else if par.right == nil {
		par.right = node
		node.parent = par
	} else { // par.right != nil
		nodeToAdd = mostLeft(par.right)
		nodeToAdd.left = node
		node.parent = nodeToAdd
	}
	h.last = node
	h.siftup()
	h.size++
}

func mostLeft(node *_node) *_node {
	// node != nil
	for node.left != nil {
		node = node.left
	}
	// node.left == nil
	return node
}

func (h *MyHeap) siftup() {
	cur := h.last

	for cur.parent != nil {
		if cur.parent.val > cur.val { // 已经符合大顶堆性质
			break
		}
		cur.parent.val, cur.val = cur.val, cur.parent.val
		cur = cur.parent
	}
}
