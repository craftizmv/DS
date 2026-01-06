package adjlist

// Graph2 - implementation containing the adj list
type Graph2 struct {
	// key
	adjList map[int][]int
}

func NewGraph2(opts ...GraphOption2) *Graph2 {
	g2 := &Graph2{adjList: map[int][]int{}}
	for _, opt := range opts {
		opt(g2)
	}

	return g2
}

type GraphOption2 func(this *Graph2)

func WithInputAdjMatrix(m map[int][]int) GraphOption2 {
	return func(this *Graph2) {
		this.adjList = m
	}
}

// Operations
// AddVertex, AddEdge

func (g2 *Graph2) AddVertex(key int) {
	// if vertex does not exist then add
	if _, ok := g2.adjList[key]; !ok {
		g2.adjList[key] = []int{}
	}
}

// assuming undirected graph
func (g2 *Graph2) AddEdge(src, dst int) {
	// assuming no need to check if already the connection exist or not
	g2.adjList[src] = append(g2.adjList[src], dst)
	g2.adjList[dst] = append(g2.adjList[dst], src)
}
