package arrays

import "errors"

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
