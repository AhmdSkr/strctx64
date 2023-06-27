package stack_test

import (
	"testing"

	"github.com/AhmdSkr/strctx64/stack"
)

var basicProvider func(uint32) stack.Basic[any]

func TestNewBasicStack(t *testing.T) {
	var datum any
	capacity := uint32(10)
	s := basicProvider(capacity)
	if err := s.Pop(); err == nil {
		t.Error("An error should be returned when popping a new (empty) stack")
	}
	if err := s.Top(&datum); err == nil {
		t.Error("An error should be returned when peeking a new (empty) stack")
	}
	if err := s.Push(1); err != nil {
		t.Error("No error should be returned when pushing to a non-full stack")
	}
}

func TestBasicStackOperations(t *testing.T) {
	var (
		datum    any    = nil
		last     any    = nil
		capacity uint32 = 10
	)

	s := basicProvider(capacity)

	for i := uint32(0); i < capacity; i++ {
		last = 1 + i%5
		if err := s.Push(last); err != nil {
			t.Error("No error should be returned when pushing to a non-full stack")
		}
		if err := s.Top(&datum); err != nil {
			t.Error("No error should be returned when peeking a non-empty stack")
		}
		if datum != last {
			t.Errorf("Top of stack should be equal to last pushed element; datum = %v, expected %v", datum, last)
		}
	}
	if err := s.Push(1); err == nil {
		t.Error("An error should be returned when pushing to a full stack")
	}

	for i := uint32(0); i < capacity; i++ {
		if err := s.Top(&datum); err != nil {
			t.Error("No error should be returned when peeking a non-empty stack")
		}
		if datum != 1+(capacity-i-1)%5 {
			t.Errorf("Top of stack should be equal to last pushed element; datum = %v, expected %v", datum, 1+(capacity-i-1)%5)
		}
		if err := s.Pop(); err != nil {
			t.Error("An error should not be returned when popping a non-empty stack")
		}
	}
	if err := s.Pop(); err == nil {
		t.Error("An error should be returned when popping an (empty) stack")
	}
	if err := s.Top(&datum); err == nil {
		t.Error("An error should be returned when peeking an (empty) stack")
	}
}
