// max sum adj sub-seq
package dp_practice

import "math"

// Reference : https://www.youtube.com/watch?v=m9-H6AUBLgY&list=PLDzeHZWIZsTomOPnCiU3J95WufjE36wsb&index=5
// Leet : https://www.naukri.com/code360/problems/maximum-sum-of-non-adjacent-elements_843261
// Intuition : For every element in the array - we can either include it or exclude it
// - meaning we have 2 choices and based on it ..
// we either choose i+1(exclude case) or i+2 (include case as we can not choose adj.)
func maxSumAdjSubSeq(input []int) int {
	if len(input) == 0 {
		return 0
	}

	// starting from 0th index.
	ans := solveAdj(input, 0)

	return ans
}

func solveAdj(input []int, idx int) int {
	if idx > len(input) {
		return 0
	}

	// reached the nth value in the array
	if idx == len(input) {
		// lets say there was only one element in the array then
		// max of all the possible sub-sequence would be that only.
		return input[0]
	}

	inclCaseEffectiveMax := input[idx] + solveAdj(input, idx+2)

	// we add 0 as we are excluding the number at current pos.
	exclCaseEffectiveMax := 0 + solveAdj(input, idx+1)

	effectiveMax := int(math.Max(float64(inclCaseEffectiveMax), float64(exclCaseEffectiveMax)))
	return effectiveMax
}
