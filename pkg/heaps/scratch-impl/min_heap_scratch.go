package scratch_impl

// Reference : https://yuminlee2.medium.com/golang-heap-data-structure-45760f9562dc

type MinHeap []int

// BuildHeap - all non-leaf nodes call SiftDown to heapify the array.
// We start with the last non-leaf node.
// Time: O(n) | Space: O(1)
func (h *MinHeap) BuildHeap(array []int) {
	lastNonLeafNodeIdx := (len(array) - 2) / 2

	for currentIdx := lastNonLeafNodeIdx; currentIdx >= 0; currentIdx-- {
		endIdx := len(array) - 1
		h.siftDown(currentIdx, endIdx) // siftDown is more efficient than siftUp
	}
}

// Time: O(logn) | Space: O(1)
// continuously swap the current node with its smaller child node until it is in the correct position
func (h *MinHeap) siftDown(currentIdx int, endIdx int) {

	leftChildIdx := currentIdx*2 + 1

	for leftChildIdx <= endIdx {
		rightChildIdx := currentIdx*2 + 2

		if rightChildIdx > endIdx {
			rightChildIdx = -1
		}

		// get the smaller child node to swap
		idxToSwap := leftChildIdx
		if rightChildIdx != -1 && (*h)[rightChildIdx] < (*h)[leftChildIdx] {
			idxToSwap = rightChildIdx
		}

		// check if value of swap node is less than the value at currentIdx
		if (*h)[idxToSwap] < (*h)[currentIdx] {
			h.swap(idxToSwap, currentIdx)
			currentIdx = idxToSwap
			leftChildIdx = currentIdx*2 + 1

		} else {
			return
		}
	}
}

// Time: O(1) | Space: O(1)
func (h MinHeap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Time: O(logn) | Space: O(1)
// continuously swap the current node with its parent node until it is in the correct position
func (h *MinHeap) siftUp() {
	currentIdx := len(*h) - 1
	parentIdx := (currentIdx - 1) / 2
	for currentIdx > 0 && (*h)[currentIdx] < (*h)[parentIdx] {
		h.swap(currentIdx, parentIdx)
		currentIdx = parentIdx
		parentIdx = (currentIdx - 1) / 2
	}
}

// Time: O(logn) | Space: O(1)
// insert a new value to the end of the tree and update heap ordering
func (h *MinHeap) Insert(value int) {
	*h = append(*h, value)
	h.siftUp()
}

// Time: O(logn) | Space: O(1)
// remove and return the minimum value and update heap ordering
func (h *MinHeap) Remove() int {
	n := len(*h)
	// swap the first element and the last element in the array
	h.swap(0, n-1)
	valueToRemove := (*h)[n-1]
	// pop the last element in the array
	*h = (*h)[:n-1]
	// call siftDown to update heap ordering
	h.siftDown(0, n-2)

	return valueToRemove
}
