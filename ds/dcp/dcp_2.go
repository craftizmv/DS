package main

import "fmt"

func modify(input []int) []int {

	if len(input) == 0 {
		return nil
	}

	//var result = make([]int, 0, len(input))
	if len(input) == 1 {
		//result = append(result, input...)
		return input
	}

	// O(n)
	var preSum = make([]int, 0, len(input))
	for i := range input {
		if i == 0 {
			preSum = append(preSum, input[i])
		} else {
			preSum = append(preSum, input[i]*preSum[i-1])
		}
	}

	fmt.Println("Presum", preSum)

	// O(n)
	var postSum = make([]int, len(input))
	for i := len(input) - 1; i >= 0; i-- {
		if i == len(input)-1 {
			postSum[i] = input[i]
		} else {
			postSum[i] = input[i] * postSum[i+1]
		}
	}

	// O(n)
	fmt.Println("Postsum", postSum)
	for i := 0; i < len(input); i++ {
		if i == 0 {
			input[i] = postSum[i+1]
		} else if i == (len(input) - 1) {
			input[i] = preSum[i-1]
		} else {
			input[i] = preSum[i-1] * postSum[i+1]
		}
	}

	return input
}

func main() {
	//given an array/slice -> o/p product of all the numbers except

	sample := []int{1}
	fmt.Println(modify(sample))

}
