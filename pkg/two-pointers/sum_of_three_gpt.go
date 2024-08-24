package two_pointers

import (
	"sort"
)

func HasTripletWithTargetSum(arr []int, target int) bool {
	// Step 1: Sort the array
	sort.Ints(arr)

	// Step 2: Iterate through the array
	for i := 0; i < len(arr)-2; i++ {
		// Initialize two pointers
		left := i + 1
		right := len(arr) - 1

		// Use two pointers to find a pair with the sum equal to target - arr[i]
		for left < right {
			sum := arr[i] + arr[left] + arr[right]
			if sum == target {
				return true
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return false
}
