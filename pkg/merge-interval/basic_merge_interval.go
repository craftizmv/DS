package merge_interval

// MergeIntervals : mergeIntervals - merges interval given in a 2D array.
// Assumption - the array is already sorted.
// Sample Data : [[1,3],[2,4],[5,6]]
func MergeIntervals(intervals [][]int) [][]int {
	// Replace this placeholder return statement with your code
	x := len(intervals)

	if x <= 1 {
		return intervals
	}

	result := make([][]int, 1)
	// adding the first interval to result
	result[0] = intervals[0]
	rLen := 1

	// iterate over the 2d metrics and compare each interval
	for i := 1; i < x; i++ {
		// interval merge.
		x1 := intervals[i][0]
		x2 := intervals[i][1]

		// current window
		w1 := result[rLen-1][0]
		w2 := result[rLen-1][1]

		// overlap is there.
		if x1 >= w1 && x1 <= w2 {
			if x2 <= w2 {
				// no need to change the window
				continue
			} else if x2 > w2 {
				// then change the result window
				result[rLen-1][1] = x2
			}
		} else if x1 > w2 {
			// create the new window in the result array
			newSlice := make([][]int, rLen+1)
			copy(newSlice, result)

			result = newSlice
			// assign the new int obj literal
			result[rLen] = []int{x1, x2}
			rLen++
		}

	}

	return result
}
