package utilities

type FIFO[T any] struct {
	Items []T
}

func (f *FIFO[T]) Push(item T) {
	f.Items = append(f.Items, item)
}

func (f *FIFO[T]) Pop() T {
	if len(f.Items) == 0 {
		panic("empty fifo")
	}

	item := f.Items[0]
	f.Items = f.Items[1:]

	return item
}

func (f *FIFO[T]) IsEmpty() bool {
	return len(f.Items) == 0
}
