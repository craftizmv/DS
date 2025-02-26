package container_based_impl

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
	old := *h      // copy by value ?
	n := len(old)  // from the new copy
	x := old[n-1]  // getting the last element.
	*h = old[:n-1] // copy rest of the elements except the last element.
	return x
}
