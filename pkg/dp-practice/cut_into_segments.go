package dp_practice

// References:
// https://www.naukri.com/code360/problems/cut-into-segments_1214651?leftPanelTabValue=PROBLEM&count=25&search=cut%20into&sort_entity=reactions_count&sort_order=DESC&customSource=studio_nav

func CutIntoSegments(n int, a int, b int, c int) int {
	// This function is a placeholder for the recursive solution.
	// The actual implementation will depend on the specific
	// problem requirements.

	maxSegmentPossible := solveRecursive(n, a, b, c)
	if maxSegmentPossible < 0 {
		// if the maxSegmentPossible is less than 0 then it means
		// we cannot cut the rod into segments of length a, b, c
		return 0
	}
	// if the maxSegmentPossible is greater than 0 then it means
	// we can cut the rod into segments of length a, b, c
	// and we can return the maxSegmentPossible
	// as the answer.

	return maxSegmentPossible
}

func solveRecursive(n, a, b, c int) int {
	// base case - if the rod is of 0 length then - how many segment possible.
	if n == 0 {
		return 0
	}

	// as (n-x) can lead to a negative value
	// thats why we need to protect it with a condition
	if n < 0 {
		// doubel check the condition - which is used to return the int min
		return -1
	}

	// sub problems - based on 3 choices
	// 1. solve(n-a) + 1 (inclusion of choosing a)
	answer1 := solveRecursive(n-a, a, b, c) + 1
	answer2 := solveRecursive(n-b, a, b, c) + 1
	answer3 := solveRecursive(n-c, a, b, c) + 1

	answer := maxOfThree(answer1, answer2, answer3)

	return answer
}

func maxOfThree(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	} else {
		return c
	}
}
