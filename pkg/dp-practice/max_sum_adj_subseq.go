// max sum adj sub-seq
package dp_practice

import "math"

// Reference : https://www.youtube.com/watch?v=m9-H6AUBLgY&list=PLDzeHZWIZsTomOPnCiU3J95WufjE36wsb&index=5
// Leet : https://www.naukri.com/code360/problems/maximum-sum-of-non-adjacent-elements_843261
// Intuition : For every element in the array - we can either include it or exclude it
// - meaning we have 2 choices and based on it ..
// we either choose i+1(exclude case) or i+2 (include case as we can not choose adj.)
// We can use recursion to solve this problem, but it will be inefficient for larger inputs.
// So, we can use memoization to store the results of subproblems and avoid redundant calculations.
// We can also use dynamic programming to solve this problem iteratively.
// MaxSumAdjSubSeq calculates the maximum sum of non-adjacent elements in the input slice.
func MaxSumAdjSubSeq(input []int) int {
	if len(input) == 0 {
		return 0
	}

	// starting from 0th index.
	// ans := solveAdj(input, 0)
	mem := make([]int, len(input))
	// Initialize memoization array with -1
	for i := 0; i < len(mem); i++ {
		mem[i] = -1
	}
	// ansWithMemoization := solveAdjWithMemoization(input, 0, mem)

	ansWithIterative := MaxSumAdjSubSeqIterative(input)

	// return ansWithMemoization
	// return ans
	return ansWithIterative
}

func solveAdj(input []int, idx int) int {
	if idx >= len(input) {
		return 0
	}

	inclCaseEffectiveMax := input[idx] + solveAdj(input, idx+2)

	// we add 0 as we are excluding the number at current pos.
	exclCaseEffectiveMax := 0 + solveAdj(input, idx+1)

	effectiveMax := int(math.Max(float64(inclCaseEffectiveMax), float64(exclCaseEffectiveMax)))
	return effectiveMax
}

// solveAdjWithMemoization is a helper function that uses
// memoization to avoid recalculating results for the same index.
func solveAdjWithMemoization(input []int, idx int, mem []int) int {
	if idx >= len(input) {
		return 0
	}

	// if we have already calculated the value for this index, return it.
	if mem[idx] != -1 {
		return mem[idx]
	}

	inclCaseEffectiveMax := input[idx] + solveAdjWithMemoization(input, idx+2, mem)

	// we add 0 as we are excluding the number at current pos.
	exclCaseEffectiveMax := 0 + solveAdjWithMemoization(input, idx+1, mem)

	effectiveMax := int(math.Max(float64(inclCaseEffectiveMax), float64(exclCaseEffectiveMax)))
	mem[idx] = effectiveMax // store the result in memoization array

	return effectiveMax
}

func MaxSumAdjSubSeqIterative(input []int) int {
	if len(input) == 0 {
		return 0
	}

	if len(input) == 1 {
		return input[0]
	}

	// Initialize dp array to store the maximum sum at each index
	dp := make([]int, len(input))
	dp[0] = input[0]
	dp[1] = int(math.Max(float64(input[0]), float64(input[1])))

	for i := 2; i < len(input); i++ {
		inclCaseEffectiveMax := input[i] + dp[i-2]
		exclCaseEffectiveMax := dp[i-1]
		dp[i] = int(math.Max(float64(inclCaseEffectiveMax), float64(exclCaseEffectiveMax)))
	}

	return dp[len(input)-1]
}

// MaxSumAdjSubSeqIterativeWithSpaceOptimization calculates the maximum sum of non-adjacent
// elements in the input slice using an iterative approach with space optimization.
// This approach reduces the space complexity from O(n) to O(1) by only keeping
// track of the last two computed values instead of the entire dp array.
// This is a common technique in dynamic programming problems where only the last two states are needed.
// It is more space-efficient than the previous iterative approach.
// This function is particularly useful for large input sizes where memory usage is a concern.
func MaxSumAdjSubSeqIterativeWithSpaceOptimization(input []int) int {
	if len(input) == 0 {
		return 0
	}

	if len(input) == 1 {
		return input[0]
	}

	prev2 := input[0]
	prev1 := int(math.Max(float64(input[0]), float64(input[1])))

	for i := 2; i < len(input); i++ {
		inclCaseEffectiveMax := input[i] + prev2
		exclCaseEffectiveMax := prev1
		current := int(math.Max(float64(inclCaseEffectiveMax), float64(exclCaseEffectiveMax)))

		prev2 = prev1
		prev1 = current
	}

	return prev1
}
