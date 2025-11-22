package dp_practice

import "math"

// Find mininum of perfect square's
// sum possible to create a target num.

func minNumOfPerfectSquares(target int) {

}

func PerfectSquareSolve(target int) int {
	// base condition
	if target == 0 {
		return 0
	}

	/// ret max as we don't want to ignore it using the min condition.
	if target < 0 {
		return math.MaxInt
	}

	min := math.MaxInt
	for i := 1; i*i <= target; i++ {
		v := PerfectSquareSolve(target-i*i) + 1
		if v < min {
			min = v
		}
	}

	return min
}
