package main

import (
	"fmt"
	"github.com/Golang-Practice/ds/graph"
	"github.com/Golang-Practice/ds/leetcode/arrays"
)

// Driver file.

func main() {
	adjMatrix := graph.InitAdjMatrixWithUserInput()
	if adjMatrix != nil {
		matrix := adjMatrix.Matrix
		for i := range matrix {
			for j := range matrix[i] {
				fmt.Printf("%d ", matrix[i][j])
			}
			fmt.Println()
		}
	}

	// Test MaxSubarrSum
	input := []int{-5, -10, 11, -8, -4, 6, 6, -6}
	fmt.Println(arrays.FindMaxSubArrSum(input))
}
