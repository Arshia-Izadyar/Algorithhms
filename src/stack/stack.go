package stack

type Node[T comparable] struct {
	Value T
	Prev  *Node[T]
}

type Stack[T comparable] struct {
	Length int
	head   *Node[T]
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{
		head:   nil,
		Length: 0,
	}
}

func (s *Stack[T]) Push(val T) {
	var node = &Node[T]{
		Value: val,
		Prev:  nil,
	}
	s.Length++
	if s.head == nil {
		s.head = node
		return
	}
	node.Prev = s.head
	s.head = node
}

func (s *Stack[T]) Pop() *T {
	s.Length = max(0, s.Length-1)
	if s.Length == 0 {
		if s.head == nil {
			return nil
		}
		var head = *s.head
		s.head = nil
		return &head.Value
	}
	var head Node[T] = *s.head
	s.head = s.head.Prev
	return &head.Value
}

func (s Stack[T]) Peak() *T {
	if s.head == nil {
		return nil
	}
	return &s.head.Value
}
