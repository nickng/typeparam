package linkedlist

// Header is a linked list header.
type Header[T any] struct{
	Head *Node[T]
	Size int
}

// Node is a generic linked list node.
type Node[T any] struct {
	Item T
	Next *Node[T]
}

// New create a new LinkedList.
func New[T any]() *Header[T] {
	return &Header[T]{Size: 0}
}

// AppendItem adds an item v to the end of the linked list.
func (h *Header[T]) AppendItem(v T) *Node[T] {
	h.Size++
	if h.Head == nil {
		h.Head = &Node[T]{Item: v}
		return h.Head
	}
	return h.Head.Append(v)
}

func (h *Header[T]) Traverse(fn func(n *Node[T])) {
	node := h.Head
	for node != nil {
		fn(node)
		node = node.Next
	}
}

// Append appends a value to the end of the current
// linked list and returns pointer to the new node.
func (n *Node[T]) Append(v T) *Node[T] {
	if n.Next == nil {
		n.Next = &Node[T]{Item: v}
		return n.Next
	}
	return n.Next.Append(v)
}
