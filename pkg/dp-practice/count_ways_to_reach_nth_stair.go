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

	// to reach the destination either we would have come from just prev stair.
	cost1 := solve(nStairs-1, cost)

	//solve(nStairs-2, cost) - will already be calculated if it is stored.
	// or we should have come from prev-2 stair.
	cost2 := solve(nStairs-2, cost)

	// In this case, we don't need to add the cost[nStairs] as nStair_th pos is the destination
	// and we can not go from there. Also, there is no cost which is given for this.
	finalCost := math.Min(float64(cost1), float64(cost2))
	return int(finalCost)
}

func CountTheCostToReachNthStairWithDp(cost []int) int {
	nStairs := len(cost)
	dp := make([]int, nStairs)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}

	// to reach the destination either we would have come from just prev stair.
	cost1 := solve2(nStairs-1, cost, dp)

	//solve(nStairs-2, cost) - will already be calculated if it is stored.
	// or we should have come from prev-2 stair.
	cost2 := solve2(nStairs-2, cost, dp)

	// In this case, we don't need to add the cost[nStairs] as nStair_th pos is the destination
	// and we can not go from there. Also, there is no cost which is given for this.
	finalCost := math.Min(float64(cost1), float64(cost2))
	return int(finalCost)
}

func CountTheCostToReachNthStairWithTabulation1D(cost []int) int {
	nStairs := len(cost)
	dp := make([]int, nStairs)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}

	// to reach the destination either we would have come from just prev stair.
	cost1 := solve3(nStairs-1, cost, dp)

	//solve(nStairs-2, cost) - will already be calculated if it is stored.
	// or we should have come from prev-2 stair.
	cost2 := dp[nStairs-2]

	// In this case, we don't need to add the cost[nStairs] as nStair_th pos is the destination
	// and we can not go from there. Also, there is no cost which is given for this.
	finalCost := math.Min(float64(cost1), float64(cost2))
	return int(finalCost)
}

func CountTheCostToReachNthStairWithTabulationSpaceOptim(cost []int) int {
	nStairs := len(cost)

	// to reach the destination either we would have come from just prev stair.
	costPrev1, costPrev2 := solve4(nStairs-1, cost)

	// In this case, we don't need to add the cost[nStairs] as nStair_th pos is the destination
	// and we can not go from there. Also, there is no cost which is given for this.
	finalCost := math.Min(float64(costPrev1), float64(costPrev2))
	return int(finalCost)
}

// solve - this is just with recursion implementation.
func solve(nStairs int, cost []int) int {
	// base case
	if nStairs == 0 {
		return cost[0]
	}

	if nStairs == 1 {
		return cost[1]
	}

	// store below results in a dp array.
	result := math.Min(float64(solve(nStairs-1, cost)), float64(solve(nStairs-2, cost))) + float64(cost[nStairs])
	return int(result)
}

// solve2 - with dp impl.
// need to take the dp array as an argument.
func solve2(nStairs int, cost []int, dp []int) int {
	// base case
	if nStairs == 0 {
		return cost[0]
	}

	if nStairs == 1 {
		return cost[1]
	}

	if val := dp[nStairs]; val > 0 {
		// value is pre-computed in DP already
		return val
	}

	// store below results in a dp array.
	dp[nStairs] = int(math.Min(float64(solve2(nStairs-1, cost, dp)), float64(solve2(nStairs-2, cost, dp))) + float64(cost[nStairs]))
	return dp[nStairs]
}

// solve3 - this is with tabulation 1D DP implementation.
func solve3(nStairs int, cost []int, dp []int) int {
	// initialize the base value in the tabulation.
	dp[0], dp[1] = cost[0], cost[1]

	// instead of recursion - should we just simply use the for loop iteration.
	// building the 1D Sp from the bottom to the top.
	for i := 2; i <= nStairs; i++ {
		dp[i] = int(math.Min(float64(dp[i-1]), float64(dp[i-2])) + float64(cost[i]))
	}

	return dp[nStairs]
}

// solve4 - this is with tabulation 1D DP implementation - space optimisation
func solve4(nStairs int, cost []int) (int, int) {
	// initialize the base value in the tabulation.
	prev1, prev2 := cost[1], cost[0]

	// instead of recursion - should we just simply use the for loop iteration.
	// building the 1D Sp from the bottom to the top.
	for i := 2; i <= nStairs; i++ {
		temp := int(math.Min(float64(prev1), float64(prev2)) + float64(cost[i]))
		prev2 = prev1
		prev1 = temp
	}

	return prev1, prev2
}
