package stack

import (
	"reflect"
	"testing"
)

func TestNewStack(t *testing.T) {
	var s = NewStack()
	var stype = reflect.TypeOf(s)

	if stype.Kind() != reflect.Ptr {
		t.Error("Invalid type returned")
	}

	if top != -1 {
		t.Errorf("Received: %d Expected: %d\n", top, -1)
	}
}

func TestIsEmpty(t *testing.T) {
	var s = NewStack()
	if s.IsEmpty() != true {
		t.Error("Failed to test when stack is empty")
	}

	s.Push(10)
	if s.IsEmpty() != false {
		t.Error("Failed to test when stack is not empty")
	}
}

func TestPush(t *testing.T) {
	var s = NewStack()
	s.Push(10)

	if top != 0 {
		t.Error("Stack size mismatch")
	}

	// Set top to maxSize to test stack overflow
	top = maxSize - 1
	s.Push(20)

	if top != maxSize-1 {
		t.Error("unable to detect stack overflow condition")
	}
}

func TestPop(t *testing.T) {
	var s = NewStack()
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Push(40)

	var popped = s.Pop()
	if popped != 40 {
		t.Error("Wrong element popped")
	}

	// Test pop when stack is empty
	// pop remaining elements too
	popped = s.Pop()
	popped = s.Pop()
	popped = s.Pop()

	popped = s.Pop()
	if popped != top {
		t.Error("Unable to empty stack")
	}
}

func TestTop(t *testing.T) {
	var s = NewStack()
	s.Push(10)
	s.Push(20)

	var sTop = s.Top()
	if sTop != 20 {
		t.Error("Wrong top element")
	}

	_ = s.Pop()
	_ = s.Pop()

	// Test top when stack is empty
	if s.Top() != -1 {
		t.Error("Failed to check if stack is empty or not")
	}
}
