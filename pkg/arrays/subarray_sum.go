package arrays

import "math"

// kadane's algo ..
func FindSubarraySum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var max, localSum = math.MinInt, 0
	for i := 0; i < len(nums); i++ {
		// check if the number is pos or neg
		num := nums[i]
		newLocalSum := localSum + num
		if newLocalSum < 0 {
			// keep local sum pos
			localSum = 0
		} else {
			localSum = newLocalSum
		}

		if newLocalSum > max {
			max = newLocalSum
		}
	}

	return max
}
