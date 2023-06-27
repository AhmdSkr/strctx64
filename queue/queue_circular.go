package queue

import (
	"fmt"
)

// Creates a static queue instance satisfying the
// [queue.Extended] interface.
func NewStatic[T any](capacity uint32) Extended[T] {
	return &circular[T]{
		front:     uint32(0),
		size:      uint32(0),
		capacity:  capacity,
		container: make([]T, capacity),
	}
}

type circular[T any] struct {
	front     uint32
	size      uint32
	capacity  uint32
	container []T
}

func (q *circular[T]) IsEmpty() bool    { return q.size == 0 }
func (q *circular[T]) IsFull() bool     { return q.size == q.capacity }
func (q *circular[T]) Size() uint32     { return q.size }
func (q *circular[T]) Capacity() uint32 { return q.capacity }
func (q *circular[T]) Enqueue(datum T) error {
	if q.IsFull() {
		return fmt.Errorf("queue is full! size = %v  have reached capacity = %v", q.Size(), q.Capacity())
	}
	index := (q.front + q.size) % q.capacity
	q.container[index] = datum
	q.size++
	return nil
}
func (q *circular[T]) Dequeue() error {
	if q.IsEmpty() {
		return fmt.Errorf("queue is empty! %v elements are in the queue", q.Size())
	}
	q.front = (q.front + 1) % q.capacity
	q.size--
	return nil
}
func (q *circular[T]) Front(dst *T) error {
	if q.IsEmpty() {
		return fmt.Errorf("queue is empty! %v elements are in the queue", q.Size())
	}
	*dst = q.container[q.front]
	return nil
}
