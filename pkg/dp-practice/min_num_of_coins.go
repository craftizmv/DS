package dp_practice

import "math"

// Reference : https://www.youtube.com/watch?v=A3FHNCAkhxE&list=PLDzeHZWIZsTomOPnCiU3J95WufjE36wsb&index=4
// leetcode : https://leetcode.com/problems/coin-change/description/

// minNumberOfCoinsToPurchase - Problem : ??
// Given denominations of coin in infinite amount - find
// the min num of coins for which a target sum can be formed.
// Normal daily - 2 use case shopkeeper ka .. usko kam se kam
// change dena pade ...
func minNumberOfCoinsToPurchase(coins []int, target int) int {

	// this condition check is valid just for the first pass
	// can be moved to the parent func call.
	if len(coins) <= 0 {
		return 0
	}

	// check if the value of target is less than 0 which means
	// this path is not containing a valid solution.
	if target < 0 {
		return -2
	}

	// as now target is reaching to a point where it is completely
	// finishing the value - it means that coin value used to subtract
	// this from previous target works.
	if target == 0 {
		// this needs to be 0 as it is previously considered to be included.
		return 0
	}

	minCoin := -1
	for i := range coins {
		resultRec := minNumberOfCoinsToPurchase(coins, target-coins[i]) + 1
		if resultRec > 0 && minCoin > 0 {
			minCoin = int(math.Min(float64(minCoin), float64(resultRec)))
		} else if resultRec > 0 && minCoin <= 0 {
			minCoin = resultRec
		}
	}

	return minCoin
}

// minNumberOfCoinsToPurchaseWithDp - with DP
func minNumberOfCoinsToPurchaseWithDp(coins []int, target int, targetDp map[int]int) int {
	// check if the value of target is less than 0 which means
	// this path is not containing a valid solution.
	if target < 0 {
		return -2
	}

	if len(coins) <= 0 {
		return 0
	}

	// as now target is reaching to a point where it is completely
	// finishing the value - it means that coin value used to subtract
	// this from previous target works.
	if target == 0 {
		// this needs to be 0 as it is previously considered to be included.
		return 0
	}

	minCoin := -1
	for i := range coins {
		remainingTarget := target - coins[i]
		var recResult int

		// leveraging DP to get the store result if present.
		if val, ok := targetDp[remainingTarget]; ok {
			recResult = val
		} else {
			recResult = minNumberOfCoinsToPurchaseWithDp(coins, target-coins[i], targetDp) + 1
		}

		if recResult > 0 && minCoin > 0 {
			minCoin = int(math.Min(float64(minCoin), float64(recResult)))
		} else if recResult > 0 && minCoin <= 0 {
			minCoin = recResult
		}

		// memoize it - store it in dp.
		targetDp[remainingTarget] = recResult

	}

	return minCoin
}

// minNumberOfCoinsToPurchaseWithTab - with tabulation implementation.
// bottom's up approach.
func minNumberOfCoinsToPurchaseWithTabulation(coins []int, target int) int {
	if len(coins) <= 0 {
		return 0
	}

	// nice se uppar wale comb build karne hain ...
	// in tabulation method also we - use the dp map
	dp := make([]int, target+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}

	// dp[i] denotes - min coins needed to form target i.
	// to reach the given target .. we need to start from bottom
	// and find the min coin for each target understanding the dependency

	// this is the base which denotes - we have 0 ways for the 0 target.
	dp[0] = 0

	// dp[i] = dp[target - coins[i]] + 1 { Note : for some numbers coins[i] > target}

	// Loop to calculate each target
	for i := 1; i <= target; i++ {
		// for each target - we need to see the combination for each coins
		for k := range coins {
			if i >= coins[k] && dp[i-coins[k]] != math.MaxInt {
				// we need to check if dp[i] was already set by another coin value.
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coins[k]]+1)))
			}
		}

	}

	if dp[target] == math.MaxInt {
		// it means that we could not end up building the required target.
		return -1
	}

	return dp[target]
}
