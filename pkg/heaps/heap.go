package heaps

// MVHeap - Heaps are generally stored as arrays.
type MVHeap []int

func (h MVHeap) add(element int) {
	h[0] = element
}
