package model

// MinHeapNode represents a node in the min heap
type MinHeapNode struct {
	// The element to be stored
	Element int

	// Index of the array from which the element is taken
	I int
}

// MinHeap represents the Min Heap data structure
type MinHeap struct {
	// Slice to store elements in heap
	Harr []MinHeapNode

	// Size of min heap
	HeapSize int
}

// NewMinHeap creates a new MinHeap with the given array and size
func NewMinHeap(mharr []MinHeapNode, size int) *MinHeap {
	heapSize := size
	harr := make([]MinHeapNode, len(mharr))

	copy(harr, mharr)
	mh := &MinHeap{
		Harr:     harr,
		HeapSize: heapSize,
	}
	i := size/2 - 1
	for i >= 0 {
		mh.MinHeapify(i)
		i--
	}
	return mh
}

// MinHeapify maintains the heap property for a subtree with root at the given index
func (heap *MinHeap) MinHeapify(i int) {
	l := heap.Left(i)
	r := heap.Right(i)
	smallest := i

	if l < heap.HeapSize && heap.Harr[l].Element < heap.Harr[i].Element {
		smallest = l
	}

	if r < heap.HeapSize && heap.Harr[r].Element < heap.Harr[smallest].Element {
		smallest = r
	}

	if smallest != i {
		heap.swap(i, smallest)
		heap.MinHeapify(smallest)
	}
}

// Left returns the index of the left child of the node at index i
func (heap *MinHeap) Left(i int) int {
	return 2*i + 1
}

// Right returns the index of the right child of the node at index i
func (heap *MinHeap) Right(i int) int {
	return 2*i + 2
}

// GetMin returns the root of the heap
func (heap *MinHeap) GetMin() MinHeapNode {
	return heap.Harr[0]
}

// ReplaceMin replaces the root with a new node x and heapifies the new root
func (heap *MinHeap) ReplaceMin(x MinHeapNode) {
	heap.Harr[0] = x
	heap.MinHeapify(0)
}

// swap swaps two elements in the heap
func (heap *MinHeap) swap(x, y int) {
	temp := heap.Harr[x]
	heap.Harr[x] = heap.Harr[y]
	heap.Harr[y] = temp
}
