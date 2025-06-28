// max sum adj sub-seq problem
// Reference: https://www.youtube.com/watch?v=m9-H6AUBLgY&list=PLDzeHZWIZsTomOPnCiU3J95WufjE36wsb&index=4
package dp_practice

import "math"

// TODO : Check tomorrow...

type CHOICE_SELECTION int

const (
	INCLUDE = iota
	EXCLUDE
)

func MaxSumAdjSubSeq(input []int) int {
	return solveRec(input, 0, 0)
}

// recursive solution - driver func will call this func.
func solveRec(input []int, idx, currSum int) int {
	// base case - if out of the array.
	if idx >= len(input) {
		// idx is out of range return 0
		return currSum
	}

	// choice 1 : include the given idx
	newSumInc := currSum + input[idx]
	newSumInc += solveRec(input, idx+2, newSumInc)

	// choice 2 : exclude this index
	newSumExc := currSum + 0
	newSumExc += solveRec(input, idx+1, newSumExc)

	effectiveMaxSumLevel := int(math.Max(float64(newSumInc), float64(newSumExc)))
	if effectiveMaxSumLevel > currSum {
		return effectiveMaxSumLevel
	} else {
		return currSum
	}
}
