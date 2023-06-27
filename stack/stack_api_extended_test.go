package stack_test

import (
	"fmt"
	"testing"

	"github.com/AhmdSkr/strctx64/stack"
)

var extendedProvider func(uint32) stack.Extended[any]

func checkStackSize(t *testing.T, s stack.Extended[any], expected uint32, context ...string) {
	if size := s.Size(); size != expected {
		var errorMessage string = fmt.Sprintf("stack.Size() = %v, expecting %v", size, expected)
		t.Error(errorMessage, ":", context)
	}
}

func checkStackCapacity(t *testing.T, s stack.Extended[any], expected uint32, context ...string) {
	if capacity := s.Capacity(); capacity != expected {
		var errorMessage string = fmt.Sprintf("stack.Capacity() = %v, expecting %v", capacity, expected)
		t.Error(errorMessage, ":", context)
	}
}

func checkStackEmpty(t *testing.T, s stack.Extended[any], expected bool, context ...string) {
	var isEmpty bool
	if isEmpty = s.IsEmpty(); isEmpty != expected {
		var errorMessage string = fmt.Sprintf("stack.IsEmpty() = %v, expecting %v", isEmpty, expected)
		t.Error(errorMessage, ":", context)
	}
	// State Relations
	// 1) (Empty stack) is equivalent to (stack's size equal to 0)
	if isEmpty {
		checkStackSize(t, s, 0)
	}
}

func checkStackFull(t *testing.T, s stack.Extended[any], expected bool, context ...string) {
	var isFull bool
	if isFull = s.IsFull(); isFull != expected {
		var errorMessage string = fmt.Sprintf("stack.IsFull() = %v, expecting %v", isFull, expected)
		t.Error(errorMessage, ":", context)
	}

	// State Relations
	// 1) (Full stack) is equivalent to (stack's size equal to stack's capacity)
	// 2) (Full stack) is equivalent to (stack is not empty)
	if isFull {
		checkStackSize(t, s, s.Capacity(), "(Full stack) should be equivalent to (stack's size equal to stack's capacity)")
		checkStackEmpty(t, s, false, "(Full stack) should be equivalent to (stack is not empty)")
	}
}

func TestNewExtendedStack(t *testing.T) {
	capacity := uint32(10)
	s := extendedProvider(capacity)

	checkStackEmpty(t, s, true)
	checkStackFull(t, s, false)
	checkStackCapacity(t, s, capacity)
	checkStackSize(t, s, 0)
}

func TestExtendedStackOperations(t *testing.T) {
	capacity := uint32(10)
	s := extendedProvider(capacity)
	s.Push(1)
	checkStackCapacity(t, s, capacity)
	checkStackEmpty(t, s, false)
	checkStackFull(t, s, false)
	checkStackSize(t, s, 1)
	s.Pop()
	checkStackCapacity(t, s, capacity)
	checkStackEmpty(t, s, true)
	checkStackFull(t, s, false)
	checkStackSize(t, s, 0)

	iterations := capacity - 1
	for i := uint32(0); i < iterations; i++ {
		s.Push(i)
		checkStackCapacity(t, s, capacity)
		checkStackEmpty(t, s, false)
		checkStackFull(t, s, false)
		checkStackSize(t, s, i+1)
	}

	s.Push(1)
	checkStackCapacity(t, s, capacity)
	checkStackEmpty(t, s, false)
	checkStackFull(t, s, true)
	checkStackSize(t, s, capacity)

	for i := uint32(0); i < iterations; i++ {
		s.Pop()
		checkStackCapacity(t, s, capacity)
		checkStackEmpty(t, s, false)
		checkStackFull(t, s, false)
		checkStackSize(t, s, iterations-i)
	}
	s.Pop()
	checkStackCapacity(t, s, capacity)
	checkStackEmpty(t, s, true)
	checkStackFull(t, s, false)
	checkStackSize(t, s, 0)
}
