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

// ThreeSumOptimisedTwoPointer : efficient way to create 3 sum.
// 1. sort and then move the 2 pointers, maintain the target in hashmap.
func ThreeSumOptimisedTwoPointer(input []int, target int) bool {
	// check the input
	if len(input) < 3 {
		return false
	}

	// sort the input array
	sort.Ints(input)

	// follow 2p approach for looking at the target
	complementLookupMap := make(map[int]struct{}, len(input)/2)
	start, end := 0, len(input)-1
	moveForward := 0
	for start < end {
		startVal := input[start]
		if moveForward == 1 {
			// check if the target is present in the map or not.
			if _, ok := complementLookupMap[startVal]; ok {
				return true
			}
		}
		endVal := input[end]
		if moveForward == -1 {
			if _, ok := complementLookupMap[endVal]; ok {
				return true
			}
		}

		complement := target - (startVal + endVal)
		complementLookupMap[complement] = struct{}{}

		// increment the 2p vars based on the target val
		if complement > target {
			start++
			moveForward = 1
		} else {
			end--
			moveForward = -1
		}
	}

	return false
}

// ThreeSumOptimisedTwoLoops : Could be using 2 loops and one hashmap we search the target.
func ThreeSumOptimisedTwoLoops(input []int, target int) ([][]int, error) {
	return nil, nil
}
