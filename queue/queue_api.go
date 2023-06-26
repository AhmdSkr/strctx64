// Package queue stores different queue implementations
package queue

// Basic interface for queues that is confined to
// FIFO data operations.
type Basic[T any] interface {
	// Enqueue stores a datum at the tail of the
	// queue.
	//
	// Enqueue returns an error if and only if
	// the queue is full, otherwise it returns
	// nil.
	EnQueue(datum T) error

	// Dequeue removes the first element from
	// the queue.
	//
	// Dequeue returns an error if and only if
	// the queue is empty, otherwise it returns
	// nil.
	DeQueue() error

	// Front sets the variiable pointed to by dst
	// equal to that stored at the front of the
	// queue.
	//
	// Front returns an error if and only if the
	// queue is empty (with no change to dst's
	// value), otherwise it returns nil.
	Front(dst *T) error
}

// Extended interface for queue that includes FIFO
// data operations with additional state queries.
type Extended[T any] interface {
	// IsEmpty returs true if the queue is empty,
	// otherwise false.
	IsEmpty() bool

	// IsFull return true if the queue has reached
	// its full capacity, otherwise false.
	IsFull() bool

	// Size returns the count of elements currently
	// stored in the queue.
	//
	// If the queue is empty then the size of the queue
	// is 0. If the queue is full then the size of the
	// queue equals its capacity.
	Size() uint32

	// Capacity returns the maximum amount of elements a
	// queue can store at once.
	Capacity() uint32
}
