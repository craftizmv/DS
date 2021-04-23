package graph

import (
	"fmt"
	"sync"
)

/***** Adj Matrix *****/

type AdjMatrix struct {
	NumOfVertex int
	Matrix      [][]int
}

func (adjMatrix *AdjMatrix) HasEdge(i, j int) int {
	return adjMatrix.Matrix[i][j]
}

func (adjMatrix *AdjMatrix) AddEdge(i, j int) {
	adjMatrix.Matrix[i][j] = 1
}

func (adjMatrix *AdjMatrix) RemoveEdge(i, j int) {
	adjMatrix.Matrix[i][j] = 0
}

func InitAdjMatrixWithUserInput() *AdjMatrix {
	var verticesCount int
	fmt.Printf("Please enter number of vertices in the Matrix ...")
	_, err := fmt.Scanf("%d", &verticesCount)
	if err != nil {
		fmt.Printf("Invalid Input")
	}

	fmt.Println("Vertex count found is ...", verticesCount)

	// Create 2D array.
	var adjMatrix *AdjMatrix
	if verticesCount > 0 {
		adjMatrix = &AdjMatrix{
			NumOfVertex: verticesCount,
		}

		adjMatrix.Matrix = make([][]int, verticesCount)
		for i := range adjMatrix.Matrix {
			adjMatrix.Matrix[i] = make([]int, verticesCount)
		}
	}

	for i := 0; i < verticesCount; i++ {
		for j := 0; j < verticesCount; j++ {
			var isEdgeConnected string
			fmt.Printf("Is edge connected between %d and %d", i, j)
			fmt.Println("Enter y or Y for yes and n or N for no.")
			fmt.Scanf("%s", &isEdgeConnected)

			if isEdgeConnected == "y" || isEdgeConnected == "y" {
				adjMatrix.Matrix[i][j] = 1
			} else {
				fmt.Println("Input is other than y or Y")
				adjMatrix.Matrix[i][j] = 0
			}
		}
	}

	return adjMatrix
}

/***** Adj List Representations *****/

// Node represents a hop in the graph.
type Node struct {
	value int
}

type AdjList struct {
	nodes []*Node // represents all the nodes - Not sure if it is needed.
	edges map[Node][]*Node
	lock  sync.RWMutex
}
