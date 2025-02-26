package main

import (
	"container/heap"
	"fmt"
	. "github.com/craftizmv/DS/pkg/heaps/container-based-impl"
)

func main() {
	// an array is not heapify yet
	arrayForInit := []int{9, 31, 40, 22, 10, 15, 1, 25, 91}
	arrayForPush := []int{9, 31, 40, 22, 10, 15, 1, 25, 91}

	// apply Init method
	minHeap := buildHeapByInit(arrayForInit)

	// apply Push method
	buildHeapByPush(arrayForPush)
	fmt.Println()

	// apply Pop method: remove and return the minimum element
	elementToRemove := heap.Pop(minHeap)
	fmt.Println("removed element by Pop method: ", elementToRemove)
	fmt.Println()

	// apply Fix method
	fmt.Println("Current min heap: ", *minHeap)

	// change number from 15 t0 50 at the index 2
	(*minHeap)[2] = 50
	fmt.Println("Modify min heap value at index 2: ", *minHeap)
	heap.Fix(minHeap, 2)

	fmt.Println("Fix min heap: ", *minHeap)
	fmt.Println()

	// apply Remove method
	fmt.Println("Current min heap: ", *minHeap)

	// remove number at index 4
	heap.Remove(minHeap, 4)
	fmt.Println("Current min heap after remove element in index 4: ", *minHeap)
}

// Time: O(n)
func buildHeapByInit(array []int) *MinHeap {
	// initialize the MinHeap that has implement the heap.Interface
	minHeap := &MinHeap{}
	*minHeap = array
	heap.Init(minHeap)

	fmt.Println("buildHeapByInit: ", *minHeap)
	return minHeap
}

// Time: O(nlogn)
func buildHeapByPush(array []int) *MinHeap {
	// initialize the MinHeap that has implement the heap.Interface
	minHeap := &MinHeap{}
	for _, num := range array {
		heap.Push(minHeap, num)
	}
	fmt.Println("buildHeapByPush: ", *minHeap)
	return minHeap
}

/* output
buildHeapByInit:  [1 10 9 22 31 15 40 25 91]
buildHeapByPush:  [1 10 9 25 22 40 15 31 91]
removed element by Pop method:  1
Current min heap:  [9 10 15 22 31 91 40 25]
Modify min heap value at index 2:  [9 10 50 22 31 91 40 25]
Fix min heap:  [9 10 40 22 31 91 50 25]
Current min heap:  [9 10 40 22 31 91 50 25]
Current min heap after remove element in index 4:  [9 10 40 22 25 91 50]
*/
