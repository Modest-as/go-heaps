package skewheap

import (
	heap "github.com/theodesp/go-heaps"
)

// Node is a leaf in the heap.
type Node struct {
	Item        heap.Item
	Right, Left *Node
}

// SkewHeap is a skew heap implementation.
type SkewHeap struct {
	Root *Node
}

func (h *SkewHeap) merge(x, y *Node) *Node {
	if x == nil {
		return y
	}

	if y == nil {
		return x
	}

	if x.Item.Compare(y.Item) == 1 {
		x, y = y, x
	}

	x.Left, x.Right = x.Right, x.Left
	x.Left = h.merge(y, x.Left)

	return x
}

// Insert adds an item into the heap.
func (h *SkewHeap) Insert(v heap.Item) heap.Item {
	h.Root = h.merge(&Node{
		Item: v,
	}, h.Root)

	return v
}

// DeleteMin deletes the minimum value and returns it.
func (h *SkewHeap) DeleteMin() heap.Item {
	v := h.Root

	h.Root = h.merge(v.Right, v.Left)

	return v.Item
}

// FindMin finds the minimum value.
func (h *SkewHeap) FindMin() heap.Item {
	return h.Root.Item
}

// Clear removes all items from the heap.
func (h *SkewHeap) Clear() {
	h.Root = nil
}
