package main

import "fmt"
import . "github.com/craftizmv/DS/pkg/heaps/scratch-impl"

func main() {
	array := []int{9, 31, 40, 22, 10, 15, 1, 25, 91}
	minHeap := &MinHeap{}

	*minHeap = array
	minHeap.BuildHeap(array)
	fmt.Println("build min heap: ", *minHeap)

	// apply Remove method
	valueToRemove := minHeap.Remove()
	fmt.Println("root value: ", valueToRemove)
	fmt.Println("min heap after Remove root: ", *minHeap)

	// apply Insert method
	// append a new value, 2
	minHeap.Insert(2)
	fmt.Println("min heap after insert value 2: ", *minHeap)
}
