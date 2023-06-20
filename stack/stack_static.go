package stack

import "fmt"

// Creates a static stack instance satisfying the
// [stack.Extended] interface.
func NewStatic[T any](capacity uint64) Extended[T] {
	return &static[T]{
		size:      uint64(0),
		capacity:  capacity,
		container: make([]T, capacity),
	}
}

type static[T any] struct {
	size      uint64
	capacity  uint64
	container []T
}

func (s *static[T]) IsEmpty() bool    { return s.size == 0 }
func (s *static[T]) IsFull() bool     { return s.size == s.capacity }
func (s *static[T]) Size() uint64     { return s.size }
func (s *static[T]) Capacity() uint64 { return s.capacity }
func (s *static[T]) Push(datum T) error {
	if s.IsFull() {
		return fmt.Errorf("stack is full! size = %v  have reached capacity = %v", s.Size(), s.Capacity())
	}
	s.container[s.size] = datum
	s.size++
	return nil
}
func (s *static[T]) Pop() error {
	if s.IsEmpty() {
		return fmt.Errorf("stack is empty! %v elements are in the stack", s.Size())
	}
	s.size--
	return nil
}
func (s *static[T]) Top(datum *T) error {
	if s.IsEmpty() {
		return fmt.Errorf("stack is empty! %v elements are in the stack", s.Size())
	}
	*datum = s.container[s.size]
	return nil
}
