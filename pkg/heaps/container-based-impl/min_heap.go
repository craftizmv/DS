package container_based_impl

// Reference : https://yuminlee2.medium.com/golang-heap-data-structure-45760f9562dc
// Google impl : https://cs.opensource.google/go/go/+/refs/tags/go1.25.5:src/container/heap/heap.go

// TODO - Get complete understanding and make video.

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Pop - this method is called after the swap of root element with the last element.
// remember any is an alias for interface {}
func (h *MinHeap) Pop() any {
	// https://chatgpt.com/share/6939110a-a9c4-8007-8a3e-12b08c0cf3f5
	old := *h      // It is a value copy of the header - containing ptr to backing array, len and capacity.
	n := len(old)  // from the new copy
	x := old[n-1]  // getting the last element.
	*h = old[:n-1] // copy rest of the elements except the last element.
	return x
}
