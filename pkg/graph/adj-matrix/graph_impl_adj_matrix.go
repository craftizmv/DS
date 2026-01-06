package adjmatrix

type Graph3 struct {
	adjMatrix [][]int
}

func NewGraph3(opts ...GraphOption3) *Graph3 {
	g3 := &Graph3{
		adjMatrix: [][]int{},
	}

	for _, opt := range opts {
		opt(g3)
	}

	return g3
}

type GraphOption3 func(this *Graph3)

// AddEdge = adds edge to the graph assuming its an undirected Graph
func (g3 *Graph3) AddEdge(src, dst int) {
	g3.adjMatrix[src][dst] = 1
	g3.adjMatrix[dst][src] = 1
}
