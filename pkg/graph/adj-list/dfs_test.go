package adjlist_test

import (
	"slices"
	"testing"

	adjlist "github.com/craftizmv/DS/pkg/graph/adj-list"
)

func TestGraph2_DFSRec(t *testing.T) {

	m := map[int][]int{
		5:  {6, 7},
		6:  {5, 8},
		7:  {5, 9},
		8:  {6, 10},
		9:  {7, 10},
		10: {8, 9},
	}

	gOpt2 := adjlist.WithInputAdjMatrix(m)

	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		opts []adjlist.GraphOption2
		// Named input parameters for target function.
		start int
		want  []int
	}{
		{
			"test adj list bfs impl",
			[]adjlist.GraphOption2{gOpt2},
			5,
			[]int{5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g2 := adjlist.NewGraph2(tt.opts...) // Unpack the slice for variadic arguments
			m := map[int]struct{}{}
			result := make([]int, 0, 10)

			m[tt.start] = struct{}{}
			result = append(result, tt.start)
			got := g2.DFSRec(tt.start, m, result)
			if !slices.Equal(got, tt.want) {
				t.Errorf("DFSRec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph2_DFSRec2(t *testing.T) {

	m := map[int][]int{
		5:  {6, 7},
		6:  {5, 8},
		7:  {5, 9},
		8:  {6, 10},
		9:  {7, 10},
		10: {8, 9},
	}

	gOpt2 := adjlist.WithInputAdjMatrix(m)

	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		opts []adjlist.GraphOption2
		// Named input parameters for target function.
		start int
		want  []int
	}{
		{
			"test adj list bfs impl",
			[]adjlist.GraphOption2{gOpt2},
			5,
			[]int{5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g2 := adjlist.NewGraph2(tt.opts...) // Unpack the slice for variadic arguments
			m := map[int]struct{}{}
			result := make([]int, 0, 10)

			m[tt.start] = struct{}{}
			result = append(result, tt.start)
			g2.DFSRec2(tt.start, m, &result)

			t.Logf("DFSRec() = %v", result)
		})
	}
}
