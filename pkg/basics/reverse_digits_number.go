package basics

import (
	"fmt"
	"math"
)

// Given an integer N - return the reverse of the given number.
// 1234 - 4321
// 10, 100
// Get the digits and then multiply by the corresponding .. 10, 100 ...etc.

// ReverseDigits - Complexity Analysis: TC : O(n), SC : O(n)
func ReverseDigits(n int) int {
	if n < 10 {
		return n
	}

	arr := make([]int, 0)
	tensCount := 0
	// < O(n)
	for {
		// base
		if n < 10 {
			arr = append(arr, n)
			break
		}
		remainder := n % 10
		arr = append(arr, remainder)
		n = n - remainder
		n /= 10
		tensCount++
	}

	fmt.Println("Arr is : ", arr, "Tens Count : ", tensCount)
	finalSum := 0
	// O(n)
	for i := 0; i < len(arr); i++ {
		if tensCount > 0 {
			finalSum += arr[i] * int(math.Pow(10, float64(tensCount)))
			tensCount--
		} else {
			finalSum += arr[i]
		}
	}

	return finalSum
}

// There can be a recursive sol as well.
