package main

import (
	"fmt"
	"strconv"
)

/*func findLowestPositive(input []int) int {

	if len(input) == 0 {
		return -1
	}

	var currentMin int
	var isSetByNeg bool

	for i := range input {
		if i == 0 {
			x := input[i]
			switch {
			case x <= 0:
				currentMin = 1
				isSetByNeg = true
			case x > 1:
				currentMin = 1
				isSetByNeg = true
			default:
				currentMin = 1
			}
		} else {
			currVal := input[i]
			switch {
			case currVal > currentMin:
				if currentMin == 1 && !isSetByNeg {
					// here we need to increment the value of the min after the current val
					// as we are moving in positive direction.
					if currVal == currentMin+1 {
						currentMin = currVal + 1
					} else {
						currentMin = currentMin + 1
					}
				}
			case currVal == currentMin:
				currentMin = currentMin + 1
			case currVal < currentMin:
				if currVal < 0 && currentMin == 1 {
					continue
				} else if currVal <= 0 && currentMin != 1 {
					currentMin = 1
					isSetByNeg = true
				} else {
					currentMin = currVal - 1
				}
			}
		}
	}

	return currentMin
}*/

func main() {
	//fmt.Println(findLowestPositive([]int{1, 2, 0}))
	oldSeason := "Shakti"
	val, err := strconv.Atoi(oldSeason[7:])
	if err != nil {
		fmt.Println("Err")
	}

	fmt.Println(val)
}
