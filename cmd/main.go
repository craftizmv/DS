package main

import (
	"fmt"
	"github.com/Golang-Practice/ds/graph"
	"github.com/Golang-Practice/ds/leetcode/arrays"
	"github.com/Golang-Practice/ds/leetcode/trees"
)

// Driver file.

func main() {

	//testAdjMatrix()

	// testMaxSubArr()

	testLevelOrderTraversal()
}

func testAdjMatrix() {
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
}

func testMaxSubArr() {
	// Test MaxSubarrSum
	input := []int{-5, -10, 11, -8, -4, 6, 6, -6}
	fmt.Println(arrays.FindMaxSubArrSum(input))
}

func testLevelOrderTraversal() {
	root := &trees.TreeNode{
		Val: 3,
		Left: &trees.TreeNode{
			Val: 9,
		},
		Right: &trees.TreeNode{
			Val: 20,
			Left: &trees.TreeNode{
				Val: 15,
			},
			Right: &trees.TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println(trees.LevelOrderUsingQueue(root))
}
