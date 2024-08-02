package twopointers

import "sort"

/*
Sum of three can be implemented by below 2 approaches.
1. Sort the array
2. Iterate the array 0 - n-1, and then maintain high and low pointers.

*/

func (input *TP) FindSumOfThree(nums []int, target int) bool {
	// Replace this placeholder return statement with your code
	if len(nums) < 3 {
		return false
	}

	// sort the slice
	sort.Ints(nums)
	lIdx := 0
	hIdx := len(nums) - 1
	for i := 0; i < len(nums)-1; i++ {
		if i == lIdx || i == hIdx {
			continue
		}
		currNum := nums[i]
		lowVal := nums[lIdx]
		highVal := nums[hIdx]
		tripSum := currNum + lowVal + highVal

		// We need to iterate on the number until we find the value we are looking for.
		if tripSum == target {
			return true
		} else if tripSum < target {
			if lowVal < highVal {
				lIdx++
			} else {
				hIdx--
			}
		} else {
			return false
		}
	}

	return false
}
