package main

import (
	"CodeGuide/base/utils"
	"fmt"
)

// 在一段大小为k的区间上，维护区间的最大值
// 初始想法是：滑动数组。添加元素没有问题，但是移除元素的时候，会影响最大值。
// 如果正好最大值被移除，那么只能遍历来寻找最大值了。
// 可不可以多维护一个 次大值，但是同样仍然有风险。因为次大值也需要更新，如果次大值被删除，需要遍历来寻找次大值。

// 有没有好方法呢？单调队列！

// tail                              top
// ------------------------------------
//                      |
//                |     |
//          |     |     |
//     |    |     |     |
// ------------------------------------
// 使用单调递减的单调队列

// *************************** 注意队列中不是严格单调，严格单调将导致错误答案 ******************
// 如果严格单调，那么相等的备选答案将不会记录到数据结构中，而由于窗口滑动导致的最佳答案移除，此时相等的备选答案本应该上位，可惜没有记录，导致上位的答案是错误的。
// 滑动区间的最大值，要注意单调队列中的单调关系的设定。
func MaxSlideWindow(nums []int, k int) []int {
	res := make([]int, len(nums)-k+1)
	queue := NewDeque()

	for i := 0; i < len(nums); i++ {
		// 先将前k-1个加入到单调队列中
		if i < k-1 {
			for !queue.IsEmpty() && utils.Less(queue.Tail(), nums[i]) {
				queue.PopBack()
			}
			// q is empty || cur <= queue.Tail
			queue.PushBack(nums[i])

			//fmt.Print(i," - ")
			//for root := queue.list.First();root!= &queue.list.root; root = root.Next {
			//	fmt.Print(root.Key, " ")
			//}
			//fmt.Println()

		} else { // 窗口开始滑动, i>= k -1
			for !queue.IsEmpty() && utils.Less(queue.Tail(), nums[i]) {
				queue.PopBack()
			}
			queue.PushBack(nums[i]) // 大小为k区间上的所有可能答案

			res[i-k+1] = queue.Front().(int)

			//fmt.Print(i, " - ")
			//for root := queue.list.First();root != &queue.list.root; root = root.Next {
			//	fmt.Print(root.Key, " ")
			//}
			//fmt.Println()

			if queue.Front().(int) == nums[i-k+1] { //若最大值是左边界，移除
				queue.PopFront()
			}
		}
	}
	return res
}

func main() {
	nums := []int{5, 3, 3, 5, 2, 3}
	fmt.Println(MaxSlideWindow(nums, 4))
}

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
//	l.PushBack(&Element{Key: 6})
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
