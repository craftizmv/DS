package tests

import (
	"errors"
)

func SmallestRange(input [][]int) ([]int, error) {
	// check if input len > 0
	if input == nil || len(input) < 1 {
		return []int{}, errors.New("input is empty")
	}

	// find the max length of subarr
	maxLen := 0
	for i := 0; i < len(input); i++ {
		if len(input[i]) > maxLen {
			maxLen = len(input[i])
		}
	}

	for idx := 0; idx < maxLen; idx++ {
		// here get all the arrays and compare the elements
		rangeArr := compareIndxAndReturnRange(idx, input)
	}

}

func compareIndxAndReturnRange(idx int, input [][]int) []int {
	result := []int{}
	for i := range input {
		subArr := input[i]
		subArrElement := subArr[idx]
		if len(result) == 0 {
			result = append(result, subArrElement)
		} else {
			if len(result) > 1 {
				// check both boundries
				if subArrElement < result[0] {
					result[0] = subArrElement
				} else if subArrElement > result[1] {
					result[1] = subArrElement
				}
			}
		}
	}

	return result
}
