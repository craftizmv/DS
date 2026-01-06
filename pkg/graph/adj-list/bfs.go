package adjlist

import "fmt"

// TODO : Correct implementation ✅ (Review logic)
// Logic : Start with a seed element (mark it pre visited) by putting item in the queue and then
// and then start a while loop to pop, add to result and then visit other neighbour nodes
// and mark them visited and push to queue (if not already visited.)
// Link to a flow chart.
func (g2 *Graph2) BFS(start int) []int {
	// temp visited array

	visited := map[int]struct{}{
		start: {},
	}

	bfs := []int{}

	// slice as queue has elements which are visited and we need to visit their neighbours
	queue := []int{start}

	// while - queue is non empty - keep on popping
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]

		// pop from the front and append to the result.
		bfs = append(bfs, front)

		// add neighbours to queue
		for idx, _ := range g2.adjList[front] {
			v := g2.adjList[front][idx]
			if _, ok := visited[v]; !ok {
				// mark visited
				visited[v] = struct{}{}
				// enqueue
				queue = append(queue, v)
			}
		}
	}

	return bfs
}

// wrong implementation : ❌
func (g2 *Graph2) BFS2(start int) []int {
	// temp visited array
	visited := map[int]struct{}{}

	bfs := []int{}

	// slice as queue has elements which are visited and we need to visit their neighbours
	queue := []int{start}

	// while - queue is non empty - keep on popping
	for len(queue) > 0 {

		// pop
		front := queue[0]
		queue = queue[1:]

		// visit - Logically wrong ❌ as u can keep pushing the item other times also
		//  .. mark visited only when it is popped it can be queue multiple times.
		visited[front] = struct{}{}

		// push to result
		bfs = append(bfs, front)

		// add neighbours to queue if not already visited
		neighbours := g2.adjList[front]
		for idx, _ := range neighbours {
			v := g2.adjList[front][idx]
			if _, ok := visited[v]; !ok {
				// enqueue
				queue = append(queue, v)
			}
		}
		fmt.Println("queue state now is : ", queue, "with front : ", front)
	}

	return bfs
}
