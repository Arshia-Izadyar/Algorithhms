package queue

import "fmt"

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type Queue[T comparable] struct {
	Length int
	Head   *Node[T]
	Tail   *Node[T]
}

func (q Queue[T]) Len() int {
	counter := 0
	for {
		counter++
		fmt.Println(q.Head.Value)
		if q.Head.Next == nil {
			break
		}
		q.Head = q.Head.Next
	}

	return counter
}

func (q Queue[T]) Peak() *T {
	if q.Head == nil {
		return nil
	}
	return &q.Head.Value
}
func (q *Queue[T]) Dequeue() *T {
	if q.Head == nil {
		return nil
	}
	// if q.Length <= 0 {
	// 	return nil
	// }
	q.Length--

	var head Node[T] = *q.Head
	head.Next = nil

	q.Head = q.Head.Next

	return &head.Value
}

func (q *Queue[T]) Enqueue(val T) {
	q.Length++
	var node = &Node[T]{Value: val, Next: nil}
	if q.Tail == nil {
		q.Head = node
		q.Tail = q.Head
	}
	q.Tail.Next = node
	q.Tail = node

}
