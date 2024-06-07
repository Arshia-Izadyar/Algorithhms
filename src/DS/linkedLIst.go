package ds

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

func (n *Node[T]) Length() int {
	var counter = 0
	for {
		counter++
		if n.Next == nil {
			break
		}
		n = n.Next

	}
	return counter
}

func Remove[T comparable](val T, n *Node[T]) *Node[T] {
	res := n
	var p Node[T]
	for {
		if res == nil {
			break
		}

		if res.Value == val {
			p.Next = (*res).Next
		}
		p = *res
		res = res.Next

	}
	return n
}

func (n *Node[T]) Append(val T) int {
	res := n
	for {
		if res.Next == nil {
			break
		}

		res = res.Next
	}

	new := Node[T]{
		Value: val,
		Next:  nil,
	}
	res.Next = &new

	return 1

}
