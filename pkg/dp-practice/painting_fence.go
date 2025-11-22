package dp_practice

// black Board :  https://onlineboard.eu/b/H4ggNTMU
// TODO- 28-4

// given n = fence, k = colors
// find = no of ways of painting such that not more than 2 cons. have
// same color

// k = constant
// n = we can keep changing to reduce the size of the problem.
func paintTheFence(n int, k int) int {
	if n == 0 || k == 0 {
		return 0
	}

	if n == 1 {
		return k
	}

	if n == 2 {
		return k * k
	}

	result := (k - 1) * (paintTheFence(n-1, k) + paintTheFence(n-2, k))
	return result
}
