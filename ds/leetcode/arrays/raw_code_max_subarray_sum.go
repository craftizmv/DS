package arrays

//kadane's algo :

// finds maximum subarray sum in an array.
func FindMaxSubArrSum(src []int) int {
	// we will solve it using the kadane's algo where we take the prev subArray sum into the account
	// If the (prev sum + num) is lesser than the num itself then we reset the currSum to the number
	// Also, we keep track of the maximum subArr sum and compare it with the currentSum. If currentSum
	// is '>' than the maxSubArrSum then update the maxSubArrSum

	// Set the currentSubArrSum and maxSubArrSum variable
	currentSubArrSum := src[0]
	maxSubArrSum := src[0]

	for i := 1; i < len(src); i++ {
		num := src[i]
		currentSubArrSum = findIntMax(num, currentSubArrSum+num)
		maxSubArrSum = findIntMax(currentSubArrSum, maxSubArrSum)
	}

	return maxSubArrSum
}

func findIntMax(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
