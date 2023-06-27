package queue_test

import (
	"fmt"
	"testing"

	"github.com/AhmdSkr/strctx64/queue"
)

var extendedProvider func(uint32) queue.Extended[any]

func checkQueueSize(t *testing.T, q queue.Extended[any], expected uint32, context ...string) {
	if size := q.Size(); size != expected {
		var errorMessage string = fmt.Sprintf("queue.Size() = %v, expecting %v", size, expected)
		t.Error(errorMessage, ":", context)
	}
}
func checkQueueCapacity(t *testing.T, q queue.Extended[any], expected uint32, context ...string) {
	if capacity := q.Capacity(); capacity != expected {
		var errorMessage string = fmt.Sprintf("queue.Capacity() = %v, expecting %v", capacity, expected)
		t.Error(errorMessage, ":", context)
	}
}
func checkQueueEmpty(t *testing.T, q queue.Extended[any], expected bool, context ...string) {
	var isEmpty bool
	if isEmpty = q.IsEmpty(); isEmpty != expected {
		var errorMessage string = fmt.Sprintf("queue.IsEmpty() = %v, expecting %v", isEmpty, expected)
		t.Error(errorMessage, ":", context)
	}
	// State Relations
	// 1) (Empty queue) is equivalent to (queue's size equal to 0)
	if isEmpty {
		checkQueueSize(t, q, 0)
	}
}
func checkQueueFull(t *testing.T, q queue.Extended[any], expected bool, context ...string) {
	var isFull bool
	if isFull = q.IsFull(); isFull != expected {
		var errorMessage string = fmt.Sprintf("queue.IsFull() = %v, expecting %v", isFull, expected)
		t.Error(errorMessage, ":", context)
	}

	// State Relations
	// 1) (Full queue) is equivalent to (queue's size equal to queue's capacity)
	// 2) (Full queue) is equivalent to (queue is not empty)
	if isFull {
		checkQueueSize(t, q, q.Capacity(), "(Full queue) should be equivalent to (queue's size equal to queue's capacity)")
		checkQueueEmpty(t, q, false, "(Full queue) should be equivalent to (queue is not empty)")
	}
}

func TestNewQueue(t *testing.T) {
	capacity := uint32(10)
	q := extendedProvider(capacity)
	checkQueueCapacity(t, q, capacity)
	checkQueueEmpty(t, q, true)
	checkQueueFull(t, q, false)
	checkQueueSize(t, q, 0)
}

func TestExtendedOperations(t *testing.T) {
	capacity := uint32(10)
	q := extendedProvider(capacity)
	q.Enqueue(1)
	checkQueueCapacity(t, q, capacity)
	checkQueueEmpty(t, q, false)
	checkQueueFull(t, q, false)
	checkQueueSize(t, q, 1)
	q.Dequeue()
	checkQueueCapacity(t, q, capacity)
	checkQueueEmpty(t, q, true)
	checkQueueFull(t, q, false)
	checkQueueSize(t, q, 0)

	iterations := capacity - 1
	for i := uint32(0); i < iterations; i++ {
		q.Enqueue(i)
		checkQueueCapacity(t, q, capacity)
		checkQueueEmpty(t, q, false)
		checkQueueFull(t, q, false)
		checkQueueSize(t, q, i+1)
	}

	q.Enqueue(1)
	checkQueueCapacity(t, q, capacity)
	checkQueueEmpty(t, q, false)
	checkQueueFull(t, q, true)
	checkQueueSize(t, q, capacity)

	for i := uint32(0); i < iterations; i++ {
		q.Dequeue()
		checkQueueCapacity(t, q, capacity)
		checkQueueEmpty(t, q, false)
		checkQueueFull(t, q, false)
		checkQueueSize(t, q, iterations-i)
	}
	q.Dequeue()
	checkQueueCapacity(t, q, capacity)
	checkQueueEmpty(t, q, true)
	checkQueueFull(t, q, false)
	checkQueueSize(t, q, 0)
}
