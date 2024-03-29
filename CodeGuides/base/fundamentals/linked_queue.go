/*
* Algorithm 4-th Edition
* Golang translation from Java by Robert Sedgewick and Kevin Wayne.
*
* @Author: Yajun
* @Date:   2020/10/31 20:25
 */

package fundamentals

import (
	"CodeGuide/base/abstract"
	"CodeGuide/base/utils"
)

// last处push，first处pop
type LinkedQueue struct {
	first *abstract.Node
	last  *abstract.Node
	n     int
}

func NewLinkedQueue() *LinkedQueue { return &LinkedQueue{} }

func (q *LinkedQueue) Enqueue(item interface{}) {
	newNode := &abstract.Node{Key: item}
	if q.last != nil {
		q.last.Next = newNode
	} else {
		q.first = newNode
	}
	q.last = newNode
	q.n++
}

func (q *LinkedQueue) Dequeue() interface{} {
	utils.Assert(!q.IsEmpty(), "queue underflow")
	item := q.first.Key
	q.first = q.first.Next
	if q.first == nil {
		q.last = q.first
	}
	q.n--
	return item
}

func (q *LinkedQueue) IsEmpty() bool { return q.first == nil }

func (q *LinkedQueue) Size() int { return q.n }

func (q *LinkedQueue) Peek() interface{} {
	utils.Assert(!q.IsEmpty(), "queue underflow")
	return q.first.Key
}

func (q *LinkedQueue) Tail() interface{} {
	utils.Assert(!q.IsEmpty(), "queue underflow")
	return q.last.Key
}

func (q *LinkedQueue) Iterate() abstract.Iterator {
	return utils.NewLinkedListIterator(q.first)
}
