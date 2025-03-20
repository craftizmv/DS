package dp_practice

import "math"

// CountWaysToNthStair - Recursive impl.
// f(k) = f(k+1) + f(k+2) where k=[0 .. n]
// we could have also used f(k) = f(k-1) + f(k-2)
// Ref : https://leetcode.com/problems/climbing-stairs/
// Ref : https://www.youtube.com/watch?v=S31W3kohFDk&list=PLDzeHZWIZsTomOPnCiU3J95WufjE36wsb&index=2

func CountWaysToNthStair(nStairs int, currentStair int) int {
	// base case 1
	// yahan pe 1 return karna hai to go from current stair to the same stair
	if currentStair == nStairs {
		return 1
	}

	// base case 2 - if current stair is greater than the total number of stairs.
	if currentStair > nStairs {
		return 0
	}

	return CountWaysToNthStair(nStairs, currentStair+1) + CountWaysToNthStair(nStairs, currentStair+2)
}

// CountTheCostToReachNthStair - Ref : https://leetcode.com/problems/climbing-stairs/
// There are 3 ways to solve it.
// 1. Recursion
// 2. Recursion + Memoization
// 3. Tabulation = Bottom Up DP
// 4. Tabulation + Space Optimisation = Bottom Up DP with space optimisation
func minCostClimbingStairs(cost []int) int {
	return CountTheCostToReachNthStair(cost)
}

func CountTheCostToReachNthStair(cost []int) int {
	nStairs := len(cost)
	cost1 := solve(nStairs-1, cost)

	//solve(nStairs-2, cost) - will already be calculated if it is stored.

	cost2 := solve(nStairs-2, cost)

	// In this we don't need to add the cost[nStairs] as nStair_th pos is the destination
	// and we can not go from there. Also, there is no cost which is given for this.
	finalCost := math.Min(float64(cost1), float64(cost2))
	return int(finalCost)
}

func solve(nStairs int, cost []int) int {
	// base case
	if nStairs == 0 || nStairs == 1 {
		return cost[nStairs]
	}

	// store below results in a dp array.
	result := math.Min(float64(solve(nStairs-1, cost)), float64(solve(nStairs-2, cost))) + float64(cost[nStairs])
	return int(result)
}

// solve2 - with dp impl.
// need to take the dp array as an argument.
func solve2(nStairs int, cost []int) int {
	// base case
	if nStairs == 0 || nStairs == 1 {
		return cost[nStairs]
	}

	// TODO : check if the result is already present in the dp array

	// store below results in a dp array.
	result := math.Min(float64(solve(nStairs-1, cost)), float64(solve(nStairs-2, cost))) + float64(cost[nStairs])
	return int(result)
}
