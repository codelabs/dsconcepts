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

// Push inserts element to the end of the stack
func (n *Node) Push(element int) {

	switch n {
	case nil:
		n.element = element
		n.next = nil
		head = n
	default:
		n.element = element
		n.next = head
		head = n
	}
}
