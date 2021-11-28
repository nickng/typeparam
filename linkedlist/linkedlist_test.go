package linkedlist

import (
	"strings"
	"testing"
)

type TSlice[T any] []T

func TestLinkedList(t *testing.T) {
	l := New[string]()
	l.AppendItem("Hello")
	l.AppendItem("World")

	var output TSlice[string]

	l.Traverse(func(n *Node[string]) {
		output = append(output, n.Item)
	})

	if want, got := "Hello, World", strings.Join(output, ", "); want != got {
		t.Fatalf("Expected %s but got %s", want, got)
	}
}
