package linkedlist

import "errors"

// Node represents one element in a doubly linked list.
type Node struct {
	Value any
	prev  *Node // Points to the previous node.
	next  *Node // Points to the next node.
}

// List maintains pointers to both ends of the list.
type List struct {
	head *Node
	tail *Node
}

// NewList creates a list and inserts all provided elements.
func NewList(elements ...any) *List {
	list := &List{}

	for _, elem := range elements {
		list.Push(elem)
	}
	return list
}

// Next returns the next node.
func (n *Node) Next() *Node {
	return n.next
}

// Prev returns the previous node.
func (n *Node) Prev() *Node {
	return n.prev
}

// Unshift inserts a new node at the beginning (head).
func (l *List) Unshift(v any) {
	node := &Node{Value: v}

	// Handle an empty list.
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	// Connect the new node before the current head.
	node.next = l.head
	l.head.prev = node
	l.head = node
}

// Push inserts a new node at the end (tail).
func (l *List) Push(v any) {
	node := &Node{Value: v}

	// Handle an empty list.
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	// Connect the new node after the current tail.
	node.prev = l.tail
	l.tail.next = node
	l.tail = node
}

// Shift removes and returns the first node's value.
func (l *List) Shift() (any, error) {
	if l.head == nil {
		return nil, errors.New("Empty list")
	}

	node := l.head
	l.head = node.next

	// Update tail if the list is now empty.
	if l.head == nil {
		l.tail = nil
	} else {
		l.head.prev = nil
	}

	return node.Value, nil
}

// Pop removes and returns the last node's value.
func (l *List) Pop() (any, error) {
	if l.tail == nil {
		return nil, errors.New("Empty list")
	}

	node := l.tail
	l.tail = node.prev

	// Update head if the list is now empty.
	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.next = nil
	}

	return node.Value, nil
}

// Reverse swaps the direction of all node connections.
func (l *List) Reverse() {
	curr := l.head

	for curr != nil {
		// Swap next and prev pointers.
		curr.next, curr.prev = curr.prev, curr.next

		// Move to the next node in the original direction.
		curr = curr.prev
	}

	// The old head becomes the new tail, and vice versa.
	l.head, l.tail = l.tail, l.head
}

// First returns the first node.
func (l *List) First() *Node {
	return l.head
}

// Last returns the last node.
func (l *List) Last() *Node {
	return l.tail
}

// Count returns the total number of nodes.
func (l *List) Count() int {
	curr := l.head
	count := 0

	// Traverse until the end of the list.
	for curr != nil {
		curr = curr.next
		count++
	}
	return count
}

// Delete removes the first node matching the given value.
func (ll *List) Delete(v any) bool {
	for curr := ll.head; curr != nil; curr = curr.next {
		if curr.Value == v {

			// Update the previous node or move the head.
			if curr.prev == nil {
				ll.head = curr.next
			} else {
				curr.prev.next = curr.next
			}

			// Update the next node or move the tail.
			if curr.next == nil {
				ll.tail = curr.prev
			} else {
				curr.next.prev = curr.prev
			}

			return true
		}
	}

	// No matching node found.
	return false
}
