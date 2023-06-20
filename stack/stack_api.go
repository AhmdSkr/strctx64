// Package stack stores different stack implementations
package stack

// Basic interface for stacks that is confined to
// LIFO data operations.
type Basic[T any] interface {
	// Push stores a datum on top of the stack.
	//
	// Push returns an error if and only if the stack is
	// full, otherwise it returns nil.
	Push(datum T) error

	// Pop removes the top element from the stack.
	//
	// Pop returns an error if and only if the stack is
	// empty, otherwise it retuns nil.
	Pop() error

	// Top sets the variable pointed to by dst equal
	// to that stored on top of the stack.
	//
	// (Simply, Top peeks the value of the top element)
	//
	// Top returns an error if and only if the stack is
	// empty (with no change to dst's value), otherwise
	// it returns nil.
	Top(dst *T) error
}

// Extended interface for stack that includes LIFO data
// operations and with additional state queries.
type Extended[T any] interface {
	Basic[T]

	// IsEmpty returns true if the stack is empty,
	// otherwise false.
	IsEmpty() bool // stack can store at once.

	// IsFull return true if the stack has reached its
	// full capacity, otherwise false.
	IsFull() bool

	// Size returns the count of elements currently
	// stored in the stack.
	//
	// If the stack is empty then the size of the stack
	// is 0. If the stack is full then the size of the
	// stack equals its capacity.
	Size() uint64

	// Capacity returns the maximum amount of elements a
	// stack can store at once.
	Capacity() uint64
}
