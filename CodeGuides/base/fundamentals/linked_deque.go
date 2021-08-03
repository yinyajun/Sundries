/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2021/8/3 14:21
 */

package fundamentals

type Deque struct {
	list *DoublyLinkedList
}

func NewDeque() *Deque { return &Deque{NewDoublyLinkedList()} }

func (q *Deque) PushFront(e interface{}) { q.list.AddFirst(&Element{Key: e}) }

func (q *Deque) PushBack(e interface{}) { q.list.AddLast(&Element{Key: e}) }

func (q *Deque) PopFront() interface{} { return q.list.DelFirst().Key }

func (q *Deque) PopBack() interface{} { return q.list.DelLast().Key }

func (q *Deque) IsEmpty() bool { return q.list.IsEmpty() }

func (q *Deque) Top() interface{} { return q.list.First().Key }

func (q *Deque) Tail() interface{} { return q.list.Last().Key }

func (q *Deque) Size() int { return q.list.size }
