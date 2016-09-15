/*
Array implementation for Stacks
*/

package stack

import "fmt"

// Maximum size of a stack
const maxSize = 100

// Stack implementation using integer slices
type Stack []int

var top int

// NewStack initiates Stack with maximum allow size
func NewStack() *Stack {
	var s = make(Stack, maxSize)
	top = -1
	return &s
}

// IsEmpty checks if stack is empty or not
func (s *Stack) IsEmpty() bool {
	if top == -1 {
		return true
	}
	return false
}

// Push inserts an element at top of the stack
func (s *Stack) Push(x int) {

	if top == maxSize-1 {
		fmt.Println("Stack overflow")
		return
	}

	var t = *s
	top++
	t[top] = x
}

// Pop removes the top element from stack and returns it
func (s *Stack) Pop() int {

	if s.IsEmpty() {
		fmt.Println("Stack is empty. Nothing to pop")
		return top
	}

	var t = *s
	var popped = t[top]
	top--

	return popped
}

// Top returns the top element of the stack
func (s *Stack) Top() int {

	if s.IsEmpty() {
		return top
	}

	var t = *s
	return t[top]
}
