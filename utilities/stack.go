package utilities

type Stack[T any] struct {
	Items []T
}

func (s *Stack[T]) Push(item T) {
	s.Items = append(s.Items, item)
}

func (s *Stack[T]) Pop() T {
	if len(s.Items) == 0 {
		panic("empty fifo")
	}

	item := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]

	return item
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.Items) == 0
}
