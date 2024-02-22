package model

type MinHeapNode struct {
	Value int
	Index int
}

type MinHeap struct {
	Nodes []MinHeapNode
}

func (h *MinHeap) Len() int {
	return len(h.Nodes)
}

func (h *MinHeap) Less(i, j int) bool {
	return h.Nodes[i].Value < h.Nodes[j].Value
}

func (h *MinHeap) Swap(i, j int) {
	h.Nodes[i], h.Nodes[j] = h.Nodes[j], h.Nodes[i]
}

func (h *MinHeap) Push(x interface{}) {
	h.Nodes = append(h.Nodes, x.(MinHeapNode))
}

func (h *MinHeap) Pop() MinHeapNode {
	old := h.Nodes
	n := len(old)
	x := old[n-1]
	h.Nodes = old[0 : n-1]
	return x
}

func (h *MinHeap) Init() {
	n := len(h.Nodes)
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
}

func (h *MinHeap) down(i, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j2 := j1 + 1
		if j2 < n && h.Less(j2, j1) {
			j1 = j2
		}
		if !h.Less(j1, i) {
			break
		}
		h.Swap(i, j1)
		i = j1
	}
}
