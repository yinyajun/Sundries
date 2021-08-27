package main

// 实现LRU
// hash链表
// 为了O(1)时间找到元素，使用hash
// 为了O(1)时间插入元素且记录插入顺序，使用双向链表
// 合起来就是hash链表

type Node struct {
	Key        interface{}
	Val        interface{}
	Next, Prev *Node
}

type DoublyLinkedList struct {
	Root Node // root节点会被修改
	Size int
}

func (l *DoublyLinkedList) lazyInit() {
	if l.Root.Next == nil {
		l.Root.Next = &l.Root
		l.Root.Prev = &l.Root
		l.Size = 0
	}
}

func NewDoublyLinkedList() *DoublyLinkedList {
	l := new(DoublyLinkedList)
	l.lazyInit()
	return l
}

func (l *DoublyLinkedList) PushBack(e *Node) {
	e.Next = &l.Root
	e.Prev = l.Root.Prev
	e.Prev.Next = e
	l.Root.Prev = e
	l.Size++
}

func (l *DoublyLinkedList) PopFront() *Node {
	if l.Size == 0 {
		return nil
	}
	e := l.Root.Next
	l.Remove(e)
	return e
}

func (l *DoublyLinkedList) Remove(e *Node) {
	e.Prev.Next = e.Next
	e.Next.Prev = e.Prev
	e.Next = nil
	e.Prev = nil
	l.Size--
}

//type LinkedHashMap struct {
//	hash map[interface{}]*Node
//	list *DoublyLinkedList
//}
//
//// 链表尾部recent，链表头部old
//func NewLinkedHashMap() *LinkedHashMap {
//	return &LinkedHashMap{
//		hash: make(map[interface{}]*Node),
//		list: NewDoublyLinkedList(),
//	}
//}
//
//func (m *LinkedHashMap) Contins(key interface{}) bool {
//	_, exist := m.hash[key]
//	return exist
//}
//
//func (m *LinkedHashMap) Get(key interface{}) (interface{}, bool) {
//	if m.Contins(key) {
//		return m.hash[key].Val, true
//	}
//	return nil, false
//}
//
//// 已经存在，更新；否则插入，且插入到尾端（recent）
//func (m *LinkedHashMap) Put(key, val interface{}) {
//	if m.Contins(key) {
//		m.hash[key].Val = val
//		return
//	}
//	newNode := &Node{key, val, nil, nil}
//	m.hash[key] = newNode
//	m.list.PushBack(newNode)
//}
//
//func (m *LinkedHashMap) Remove(key interface{}) {
//	if !m.Contins(key) {
//		return
//	}
//	m.list.Remove(m.hash[key])
//	delete(m.hash, key)
//}
//
//// 移除最老的
//func (m *LinkedHashMap) PopFront() {
//	e := m.list.PopFront()
//	if e == nil { // size ==0
//		return
//	}
//	delete(m.hash, e.Key)
//}

type LRU struct {
	list *DoublyLinkedList
	hash map[interface{}]*Node
	cap  int
}

func NewLRU(cap int) *LRU {
	return &LRU{NewDoublyLinkedList(), make(map[interface{}]*Node), cap}
}

// 不存在，拉倒；存在，移动到队尾
func (c *LRU) Get(key interface{}) (interface{}, bool) {
	node, ok := c.hash[key]
	if !ok {
		return nil, false
	}
	c.list.Remove(node)
	c.list.PushBack(node)
	// todo: c.hash update?
	return node.Val, true
}

// 已经存在，更新并移动到队尾；
// 不存在，判断容量；容量满了，移除最旧的；从队尾插入新节点
func (c *LRU) Put(key, val interface{}) {
	node, ok := c.hash[key]
	if ok {
		node.Val = val
		c.list.Remove(node)
		c.list.PushBack(node)
		return
	}
	// key不存在
	if c.list.Size == c.cap {
		node = c.list.PopFront()
		delete(c.hash, node.Key)
	}
	node = &Node{key, val, nil, nil}
	c.list.PushBack(node)
	c.hash[key] = node
}

//func main() {
//	cache := NewLRU(2)
//
//	cache.Put(1, 1)
//	cache.Put(2, 2)
//	for cur := cache.list.Root.Next; cur != &cache.list.Root; cur = cur.Next {
//		fmt.Printf("[%d, %d] ", cur.Key, cur.Val)
//	}
//	fmt.Println()
//
//	fmt.Println(cache.Get(1))
//	for cur := cache.list.Root.Next; cur != &cache.list.Root; cur = cur.Next {
//		fmt.Printf("[%d, %d] ", cur.Key, cur.Val)
//	}
//	fmt.Println()
//
//	cache.Put(3, 3)
//	for cur := cache.list.Root.Next; cur != &cache.list.Root; cur = cur.Next {
//		fmt.Printf("[%d, %d] ", cur.Key, cur.Val)
//	}
//	fmt.Println()
//
//	fmt.Println(cache.Get(2))
//	for cur := cache.list.Root.Next; cur != &cache.list.Root; cur = cur.Next {
//		fmt.Printf("[%d, %d] ", cur.Key, cur.Val)
//	}
//	fmt.Println()
//
//	cache.Put(1, 4)
//	for cur := cache.list.Root.Next; cur != &cache.list.Root; cur = cur.Next {
//		fmt.Printf("[%d, %d] ", cur.Key, cur.Val)
//	}
//	fmt.Println()
//}
