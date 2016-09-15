package ds

// Node represets an element in a stack where it contains
// integer element and pointer to next node
type Node struct {
	element int
	next    *Node
}

var head *Node

// NewStack initializes stack where head pointing to nil
func NewStack() *Node {
	head = nil
	return head
}
