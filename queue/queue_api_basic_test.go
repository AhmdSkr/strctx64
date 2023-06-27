package queue_test

import (
	"testing"

	"github.com/AhmdSkr/strctx64/queue"
)

var basicProvider func(uint32) queue.Basic[any]

func TestNewBasicQueue(t *testing.T) {
	var datum any
	capacity := uint32(10)
	q := basicProvider(capacity)
	if err := q.Dequeue(); err == nil {
		t.Error("An error should be returned when dequeueing a new (empty) queue")
	}
	if err := q.Front(&datum); err == nil {
		t.Error("An error should be returned when peeking a new (empty) queue")
	}
	if err := q.Enqueue(1); err != nil {
		t.Error("No error should be returned when enqueueing to a non-full queue")
	}
}

func TestBasicQueueOperations(t *testing.T) {
	var (
		datum    any    = nil
		first    any    = nil
		capacity uint32 = 10
	)
	q := basicProvider(capacity)

	for i := uint32(0); i < capacity; i++ {
		if first == nil {
			first = 1 + i%5
			if err := q.Enqueue(first); err != nil {
				t.Error("No error should be returned when enqueueing to a non-full queue")
			}
		} else {
			if err := q.Enqueue(1 + i%5); err != nil {
				t.Error("No error should be returned when enqueueing to a non-full queue")
			}
		}
		if err := q.Front(&datum); err != nil {
			t.Error("No error should be returned when peeking a non-empty queue")
		}
		if datum != first {
			t.Errorf("Front of queue should be equal to first enqueueed element; datum = %v, expected %v", datum, first)
		}
	}
	if err := q.Enqueue(1); err == nil {
		t.Error("An error should be returned when enqueueing to a full queue")
	}

	for i := uint32(0); i < capacity; i++ {
		if err := q.Front(&datum); err != nil {
			t.Error("No error should be returned when peeking a non-empty queue")
		}
		if datum != 1+i%5 {
			t.Errorf("Front of queue should be equal to first enqueueed element; datum = %v, expected %v", datum, 1+i%5)
		}
		if err := q.Dequeue(); err != nil {
			t.Error("An error should not be returned when dequeueing a non-empty queue")
		}
	}
	if err := q.Dequeue(); err == nil {
		t.Error("An error should be returned when dequeueing an (empty) queue")
	}
	if err := q.Front(&datum); err == nil {
		t.Error("An error should be returned when peeking an (empty) queue")
	}
}
