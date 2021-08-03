package main

import (
	"CodeGuide/base/utils"
)

// 在一段大小为k的区间上，维护区间的最大值
// 初始想法是：滑动数组。添加元素没有问题，但是移除元素的时候，会影响最大值。
// 如果正好最大值被移除，那么只能遍历来寻找最大值了。
// 可不可以多维护一个 次大值，但是同样仍然有风险。因为次大值也会被删除，有需要遍历来寻找次大值。

// 有没有好方法呢？单调队列！

// tail                              top
// ------------------------------------
//                      |
//                |     |
//          |     |     |
//     |    |     |     |
// ------------------------------------
// 使用单调递减的单调队列

func MaxSlideWindow(nums []int, k int) []int {
	res := make([]int, len(nums)-k+1)
	queue := NewDeque()

	for i := 0; i < len(nums); i++ {
		// 先将前k-1个加入到单调队列中
		if i < k-1 {
			for !queue.IsEmpty() && !utils.Less(nums[i], queue.Tail()) {
				queue.PopBack()
			}
			// q is empty || cur < queue.Tail
			queue.PushBack(nums[i])
		} else { // 窗口开始滑动, i>= k -1
			for !queue.IsEmpty() && !utils.Less(nums[i], queue.Tail()) {
				queue.PopBack()
			}
			queue.PushBack(nums[i])
			res[i-k+1] = queue.Front().(int)
			if queue.Front().(int) == nums[i-k+1] {
				queue.PopFront()
			}
		}
	}
	return res
}

//func main() {
//	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
//	fmt.Println(MaxSlideWindow(nums, 3))
//}

type Element struct {
	Key, Value interface{}
	Prev, Next *Element
}

type DoublyLinkedList struct {
	root Element
	size int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return new(DoublyLinkedList).Init()
}

func (l *DoublyLinkedList) Init() *DoublyLinkedList {
	l.root.Next = &l.root
	l.root.Prev = &l.root
	l.size = 0
	return l
}

func (l *DoublyLinkedList) AddLast(e *Element) {
	e.Next = &l.root
	e.Prev = l.root.Prev
	l.root.Prev.Next = e
	l.root.Prev = e
	l.size++
}

func (l *DoublyLinkedList) AddFirst(e *Element) {
	e.Next = l.root.Next
	e.Prev = &l.root
	l.root.Next.Prev = e
	l.root.Next = e
	l.size++
}

func (l *DoublyLinkedList) del(e *Element) *Element {
	if l.IsEmpty() {
		panic("underflow")
	}
	e.Prev.Next = e.Next
	e.Next.Prev = e.Prev
	e.Next, e.Prev = nil, nil
	l.size--
	return e
}

func (l *DoublyLinkedList) DelFirst() *Element { return l.del(l.root.Next) }

func (l *DoublyLinkedList) DelLast() *Element { return l.del(l.root.Prev) }

func (l *DoublyLinkedList) IsEmpty() bool { return l.size == 0 }

func (l *DoublyLinkedList) First() *Element { return l.root.Next }

func (l *DoublyLinkedList) Last() *Element { return l.root.Prev }

type Deque struct {
	list *DoublyLinkedList
}

func NewDeque() *Deque { return &Deque{NewDoublyLinkedList()} }

func (q *Deque) PushFront(e interface{}) {
	q.list.AddFirst(&Element{Key: e})
}

func (q *Deque) PushBack(e interface{}) { q.list.AddLast(&Element{Key: e}) }

func (q *Deque) PopFront() interface{} { return q.list.DelFirst().Key }

func (q *Deque) PopBack() interface{} { return q.list.DelLast().Key }

func (q *Deque) IsEmpty() bool { return q.list.IsEmpty() }

func (q *Deque) Front() interface{} { return q.list.First().Key }
func (q *Deque) Tail() interface{}  { return q.list.Last().Key }

//func main() {
//	l := NewDoublyLinkedList()
//	l.AddFirst(&Element{Key: 5})
//	l.AddLast(&Element{Key: 6})
//	l.AddFirst(&Element{Key: 7})
//	fmt.Println(l.First())
//	fmt.Println(l.Last())
//	q := NewDeque()
//	q.PushFront(1)
//	q.PushFront(2)
//	q.PushFront(3)
//	q.PushFront(4)
//	q.PushBack(5)
//	fmt.Println(q.PopBack())
//	fmt.Println(q.PopFront())
//}
