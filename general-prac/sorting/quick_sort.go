package sorting

import "fmt"

// Revise Resources Links :
// https://www.interviewbit.com/tutorial/quicksort-algorithm/#quicksort-algorithm
// https://www.youtube.com/watch?v=Z8svOqamag8&t=65&ab_channel=KunalKushwaha

// QuickSort Implements the sorting of the slice.
// low - lowerIndex of the 'input' arr
// high - higher Index of the 'input' arr
func QuickSort(input *[]int, low, high int) {

	if input == nil {
		return
	}

	// 0. Base case if low index is equal to high or greater.
	if low >= high {
		return
	}

	// 1. First we need to find the pivot by partinitioning the array
	pivot := partitionArr(input, low, high)

	// 2. Recursively quick sort on the left
	QuickSort(input, low, pivot-1)

	// 3. Recursively quick sort on the right
	QuickSort(input, pivot+1, high)

	fmt.Println(input)

}

// Logic for partitioning the arr. choosing the right most element as pivot
// first.
func partitionArr(input *[]int, low, high int) int {

	// assign lower index to pivot
	pivot := low

	// iterate through the array to see the value of the elements
	for i := low; i < high; i++ {

		// if pivot value is greater than the index value
		if (*input)[high] > (*input)[i] {
			// here we need to perform the swap between the pivot and low element
			swap(input, pivot, i)

			// increment pivot index
			pivot++
		}
	}

	// Here we need to swap .. to put the rightmost pivot to intended position.
	swap(input, pivot, high)

	// return the pivot index
	return pivot
}

func swap(input *[]int, low, high int) {

	if low >= high {
		return
	}

	temp := (*input)[high]
	(*input)[high] = (*input)[low]
	(*input)[low] = temp
}
