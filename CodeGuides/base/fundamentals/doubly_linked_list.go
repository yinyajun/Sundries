/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2021/8/3 14:21
 */

package fundamentals

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
