package ds

import "fmt"

type N[T comparable] struct {
	Value T
	Next  *N[T]
	Prev  *N[T]
}

type DoublyLinkedList[T comparable] struct {
	Len  int
	Head *N[T]
	Tail *N[T]
}

func NewDoublyLinkedList[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		Len:  0,
		Head: nil,
	}
}

func (d *DoublyLinkedList[T]) Prepend(val T) {
	var n = &N[T]{
		Value: val,
		Next:  nil,
	}
	d.Len++
	if d.Head == nil {
		d.Head = n
		d.Tail = n
		return
	}

	n.Next = d.Head
	d.Head.Prev = n
	d.Head = n
}

func (d *DoublyLinkedList[T]) InsertAt(val T, idx int) {

	if idx > d.Len {
		fmt.Println("out of bound")
		return
	}
	if idx == d.Len {
		d.Append(val)
		return
	}
	if idx == 0 {
		d.Prepend(val)
	}

	curr := d.Head

	for i := 0; i < idx; i++ {
		curr = curr.Next
	}

	node := &N[T]{
		Value: val,
	}

	d.Len++
	node.Next = curr
	node.Prev = curr.Prev

	curr.Prev.Next = node
	curr.Prev = node

	// fmt.Println(node.Next.Value)

}

func (d *DoublyLinkedList[T]) Append(val T) {
	d.Len++
	node := &N[T]{
		Value: val,
	}
	if d.Tail == nil {
		d.Head = node
		d.Tail = node
		return
	}

	node.Prev = d.Tail
	node.Next = nil
	d.Tail.Next = node
	d.Tail = node

}

func (d *DoublyLinkedList[T]) Remove(val T) *T {
	curr := d.Head
	len := d.Len
	for i := 0; i < len; i++ {
		if curr.Value == val {
			break

		}
		curr = curr.Next
	}
	if curr == nil {
		return nil
	}
	d.Len--
	if d.Len == 0 {
		d.Head = nil
		d.Tail = nil
		return &curr.Value
	}
	value := curr.Value
	if curr == d.Head {
		d.Head = curr.Next
		return &value
	}
	if curr == d.Tail {
		d.Head = curr.Prev
		return &value
	}

	curr.Prev.Next = curr.Next
	curr.Next.Prev = curr.Prev
	curr.Next = nil
	curr.Prev = nil
	return &value
}

func (d *DoublyLinkedList[T]) GetAt(idx int) *T {
	curr := d.Head

	for i := 0; i < idx; i++ {
		if curr != nil {
			curr = curr.Next
		} else {
			break
		}
	}
	if curr == nil {
		return nil
	}
	return &curr.Value
}

func (d *DoublyLinkedList[T]) RemoveAt(idx int) {
	curr := d.Head

	for i := 0; i < idx; i++ {
		if curr != nil {
			curr = curr.Next
		} else {
			break
		}
	}
	if curr == nil {
		return
	}
	d.Len--

	if curr == d.Tail {
		d.Tail = curr.Prev
		return
	}
	if curr == d.Head {
		d.Head = curr.Next
		return
	}

	curr.Prev.Next = curr.Next

	curr.Next.Prev = curr.Prev
	curr.Next = nil
	curr.Prev = nil

}
