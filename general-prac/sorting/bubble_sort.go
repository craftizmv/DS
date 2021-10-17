package sorting

// Revise links for bubble sort : https://www.interviewbit.com/tutorial/bubble-sort/#bubble-sort

// BubbleSort sort is aslo an in-place sorting algorithm
func BubbleSort(input *[]int) *[]int {
	if input == nil {
		return input
	}

	for i := 0; i < len(*input); i++ {
		isSwapped := false

		x := len(*input) - i - 1
		for j := 0; j < x; j++ {
			if (*input)[j] > (*input)[j+1] {
				// here we need to swap
				swap(input, j, j+1)
				isSwapped = true
			}
		}

		if !isSwapped {
			break
		}
	}

	return input
}

// does modification to the input.
func swap2(input *[]int, x, y int) {
	temp := (*input)[y]
	(*input)[y] = (*input)[x]
	(*input)[x] = temp
}
