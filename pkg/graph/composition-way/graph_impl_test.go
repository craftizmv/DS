package compositionway

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Graph(t *testing.T) {
	var adjacencyList = map[int][]int{
		1: {2, 4},
		2: {3, 5, 1},
		3: {6, 2},
		4: {1, 5, 7},
		5: {2, 6, 8, 4},
		6: {3, 0, 9, 5},
		7: {4, 8},
		8: {5, 9, 7},
		9: {6, 0, 8},
	}

	g := NewGraph(WithAdjacencyList(adjacencyList))

	for key, list := range adjacencyList {
		require.ElementsMatch(t, list, g.Neighbors(key))
	}
}
