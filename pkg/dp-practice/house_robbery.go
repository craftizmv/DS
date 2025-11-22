// house robbery problem
// This is a classic dynamic programming problem where you want to maximize the amount of money you can
// rob from a series of houses, given that you cannot rob two adjacent houses.
// The problem can be solved using a recursive approach with memoization or a dynamic programming approach.
// The idea is to either rob the current house and skip the next one, or skip the current house and consider the next one.
// Also, there is catch to this problem that the first and last house are also adjacent, so you need to handle that case as well.

// Reference: https://www.youtube.com/watch?v=Fe2GeXEzWM0&list=PLDzeHZWIZsTomOPnCiU3J95WufjE36wsb&index=5
// Problem Ref : https://www.naukri.com/code360/problems/house-robber-ii_839733
package dp_practice

// HouseRobbery func calculates the maximum amount of money that can be robbed from a series of houses.
// It uses a dynamic programming approach with tabulation and space optimization.
func HouseRobbery(houses []int) int {
	if len(houses) == 0 {
		return 0
	}
	if len(houses) == 1 {
		return houses[0]
	}

	// We have two cases:
	// 1. Rob the first house and skip the last house.
	// 2. Skip the first house and rob the last house.
	return max(robHousesTabulationAndSpaceOptim(houses, 0, len(houses)-2), robHousesTabulationAndSpaceOptim(houses, 1, len(houses)-1))
}

// robHousesTabulationAndSpaceOptim is a helper function that calculates the maximum amount of money that can be robbed
// from a range of houses defined by startIdx and endIdx.
func robHousesTabulationAndSpaceOptim(houses []int, startIdx, endIdx int) int {
	if startIdx > endIdx {
		return 0
	}

	// If there's only one house in the range, return its value.
	if startIdx == endIdx {
		return houses[startIdx]
	}

	prev1, prev2 := 0, 0
	for i := startIdx; i <= endIdx; i++ {
		// Calculate the maximum amount of money that can be robbed
		// by either robbing the current house (houses[i]) and adding it to the
		// amount robbed from two houses before (prev2), or skipping the current house
		// and taking the amount robbed from the previous house (prev1).
		current := max(prev1, prev2+houses[i])
		// Update the previous amounts for the next iteration.
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

func robHouseRecursive(houses []int, idx int) int {
	if idx >= len(houses) {
		return 0
	}

	// Include the current house and skip the next one.
	includeCurrent := houses[idx] + robHouseRecursive(houses, idx+2)

	// Skip the current house and consider the next one.
	skipCurrent := robHouseRecursive(houses, idx+1)

	// Return the maximum of both choices.
	return max(includeCurrent, skipCurrent)
}

func robHouseRecursiveWithMemoization(houses []int, idx int, memo []int) int {
	if idx >= len(houses) {
		return 0
	}

	// If we have already calculated the value for this index, return it.
	if memo[idx] != -1 {
		return memo[idx]
	}

	// Include the current house and skip the next one.
	includeCurrent := houses[idx] + robHouseRecursiveWithMemoization(houses, idx+2, memo)

	// Skip the current house and consider the next one.
	skipCurrent := robHouseRecursiveWithMemoization(houses, idx+1, memo)

	// Store the maximum of both choices in the memo array.
	memo[idx] = max(includeCurrent, skipCurrent)
	return memo[idx]
}
