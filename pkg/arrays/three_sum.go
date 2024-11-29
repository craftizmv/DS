package arrays

import (
	"errors"
	"sort"
)

func ThreeSumBruteForce(input []int, target int) (bool, error) {
	if len(input) < 3 {
		return false, errors.New("invalid input")
	}

	// Intuition : Form all the combination of numbers and check if the get to a sum.
	for i := 0; i < len(input)-2; i++ {
		// J is going till last
		x := input[i]
		for j := i + 1; j < len(input)-1; j++ {
			// K is going till last
			y := input[j]
			for k := j + 1; k < len(input); k++ {
				currSum := x + y + input[k]
				if currSum == target {
					return true, nil
				}
			}
		}
	}

	return false, errors.New("not found")
}

// ThreeSumBruteForceTriplet - here the only requirement is to return the array of array of the found triplets.
// Complexity : O(n^3)
func ThreeSumBruteForceTriplet(input []int, target int) ([][]int, error) {
	var result [][]int
	if len(input) < 3 {
		return result, errors.New("invalid input")
	}

	// Intuition : Form all the combination of numbers and check if the get to a sum.
	for i := 0; i < len(input)-2; i++ {
		// J is going till last
		x := input[i]
		for j := i + 1; j < len(input)-1; j++ {
			// K is going till last
			y := input[j]
			for k := j + 1; k < len(input); k++ {
				z := input[k]
				currSum := x + y + z
				if currSum == target {
					// Add to the result array
					result = append(result, []int{x, y, z})
				}
			}
		}
	}

	if len(result) == 0 {
		return nil, errors.New("not found")
	}

	return result, nil
}

func ThreeSumTwoPointerTriplets(input []int, target int) [][]int {
	var result [][]int
	if len(input) < 3 {
		return result
	}

	sort.Ints(input)

	// we will navigate the outer loop until the n-2
	// we will run the inner loop - while l < r
	for i := 0; i < len(input)-2; i++ {
		x := input[i]
		l := i + 1
		r := len(input) - 1

		if i > 0 && input[i] == input[i-1] {
			continue
		}
		for l < r {
			y := input[l]
			z := input[r]
			currSum := x + y + z
			if currSum == target {
				result = append(result, []int{x, y, z})
				// Skip consecutive duplicates for `left` and `right`
				for l < r && input[l] == input[l+1] {
					l++
				}
				for l < r && input[r] == input[r-1] {
					r--
				}
				l++
				r--
			} else if target > 0 {
				r--
			} else {
				l++
			}
		}
	}

	return result
}

// ThreeSumOptimisedTwoPointer : efficient way to create 3 sum.
// 1. sort and then move the 2 pointers, maintain the target in hashmap.

// ThreeSumOptimisedTwoLoops : Could be using 2 loops and one hashmap we search the target.
func ThreeSumOptimisedTwoLoops(input []int, target int) ([][]int, error) {
	return nil, nil
}
