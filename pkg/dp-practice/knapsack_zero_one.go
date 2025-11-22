package dp_practice

import (
	"math"
)

// prepare the board for this
// solve :
// ref prob : https://www.naukri.com/code360/problems/0-1-knapsack_920542?interviewProblemRedirection=true&search=knapsa&count=25&sort_entity=order&sort_order=ASC&leftPanelTabValue=PROBLEM&customSource=studio_nav
// ref vid : https://youtu.be/xdPv2SZJLVI?si=z6nqPotY8ghkv9RM
// Return : Max value a person theif can return

// ToDo:
// WATCH : https://www.youtube.com/watch?v=GqOmJHQZivw (striver)

// return the max weight the theif can steal
func Knapsack(w []int, v []int, tC int) int {
	if len(w) == 0 || len(v) == 0 {
		return 0
	}

	// calling helper method.
	return solve(w, v, tC, 0)
}

func KnapsackMem(w []int, v []int, tC int, dp [][]int) int {
	if len(w) == 0 || len(v) == 0 {
		return 0
	}

	// calling helper method.
	return solveMem(w, v, 0, tC, dp)
}

func KnapsackTab(w []int, v []int, tC int) int {
	if len(w) == 0 || len(v) == 0 {
		return 0
	}

	// calling helper method.
	return solveTab(w, v, tC)
}

func KnapsackTab2(w []int, v []int, tC int) int {
	if len(w) == 0 || len(v) == 0 {
		return 0
	}

	// calling helper method.
	return solveTab2(w, v, tC)
}

// tC - target capacity
// w = all weights
// v = all volumes
// pure recursive solution
func solve(w []int, v []int, tC, i int) int {
	// bc - 2 - if my knapsack is filled.

	// we can not return 0 here but need to return min value as we do not want to include
	// this value in the prev call as the value was already added in the rec call.
	// other option is the put this as pre-condition before even calling the func so that this
	// call is ignored.
	if tC < 0 {
		return math.MinInt32
	}

	if tC == 0 {
		// if target is reached then return 0
		return 0
	}

	// bc 1 - surpassed valid array indexes
	if i >= len(w) {
		return 0
	}

	// apply the include exclude pattern
	inc := v[i] + solve(w, v, tC-w[i], i+1)
	exc := 0 + solve(w, v, tC, i+1)

	return int(math.Max(float64(inc), float64(exc)))
}

// func with memoization to reduce the unnecessary compute.
// takes in dp array with cap as tC * idx
func solveMem(w []int, v []int, idx, tC int, dp [][]int) int {
	if tC == 0 {
		// if target is reached then return 0
		return 0
	}

	// bc 1 - surpassed valid array indexes
	if idx >= len(w) {
		return 0
	}

	if dp[idx][tC] != -1 {
		// dp[idx][tC] - stores the result of the recursive call solveMem(idx, tC)
		return dp[idx][tC]
	}

	// apply the include exclude pattern
	exc := 0 + solveMem(w, v, idx+1, tC, dp)

	var inc int
	if tC >= w[idx] {
		inc = v[idx] + solveMem(w, v, idx+1, tC-w[idx], dp)
	}

	dp[idx][tC] = max(exc, inc)
	return dp[idx][tC]
}

// Refer : https://miro.com/app/board/uXjVJxLK6ZA=/?share_link_id=754504179955
// Refer the approach 2 given in the above link.
func solveTab(w, v []int, tC int) int {
	// 1. create the dp array for tab and initialize it with the base conditions.
	// b.c = when cap becomes <=0 or we reach the max given weight array length.
	n := len(w)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, tC+1)
	}

	// as in the original rec call we are making i go foreward to reach and hit the wall.
	// in that way the dependency flow is from  0 -> 1 .... -> n-3 -> n-2 -> n-1 -> n
	// dp[i][w] = max weight which can be filled from i-th position till the end
	// ∀ i = [s = n-1, end = 0] and w = [s = 0, end = tC (which is basically the max capacity.)]

	for i := n - 1; i >= 0; i-- {
		w_i := w[i]
		v_i := v[i]

		for cap := 0; cap <= tC; cap++ {
			// here we need to consider include and exclude case
			// apply the include exclude pattern
			exc := 0 + dp[i+1][cap]

			var inc int
			if cap >= w_i {
				inc = v_i + dp[i+1][cap-w_i]
			}

			dp[i][cap] = max(exc, inc)
		}
	}

	return dp[0][tC]

}

// TODO : Debug with a small example.
func solveTab2(w, v []int, tC int) int {
	// 1. create the dp array for tab and initialize it with the base conditions.
	// b.c = when cap becomes <=0 or we reach the max given weight array length.
	n := len(w)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, tC+1)
	}

	// as in the original rec call we are making i go foreward to reach and hit the wall.
	// in that way the dependency flow is from  0 -> 1 .... -> n-3 -> n-2 -> n-1 -> n
	// dp[i][w] = max weight which can be filled from i-th position till the end
	// ∀ i = [s = n-1, end = 0] and w = [s = 0, end = tC (which is basically the max capacity.)]

	for i := n - 1; i >= 0; i-- {
		w_i := w[i]
		v_i := v[i]

		// iterating backwards from the capacity perspective.
		for cap := tC; cap >= 0; cap-- {
			// here we need to consider include and exclude case
			// apply the include exclude pattern
			exc := 0 + dp[i+1][cap]

			var inc int
			if cap >= w_i {
				inc = v_i + dp[i+1][cap-w_i]
			}

			dp[i][cap] = max(exc, inc)
		}
	}

	return dp[0][tC]

}
