package adjlist

import "fmt"

func (g2 *Graph2) DFSRec(start int, visited map[int]struct{}, result []int) []int {
	for i := range g2.adjList[start] {
		num := g2.adjList[start][i]
		if _, ok := visited[num]; !ok {
			// if not visited then call the recursive dfs func.
			visited[num] = struct{}{}
			result = append(result, num)
			// we need to store the result in slice - as header
			// will keep changing on re-allocations of objects.
			// it needs to be ``` result = g2.DFSRec(num, visited, result) ```
			// else it will not work. or you need to point a slice struct which contains the
			// backing array - header address, capacity and the length.
			result = g2.DFSRec(num, visited, result)
		}
	}

	fmt.Println("result is : ", result, "visited : ", visited)

	return result
}

func (g2 *Graph2) DFSRec2(start int, visited map[int]struct{}, result *[]int) {
	for i := range g2.adjList[start] {
		num := g2.adjList[start][i]
		if _, ok := visited[num]; !ok {
			// if not visited then call the recursive dfs func.
			visited[num] = struct{}{}
			*result = append(*result, num)
			// we need to store the result in slice - as header
			// will keep changing on re-allocations of objects.
			// it needs to be ``` result = g2.DFSRec(num, visited, result) ```
			// else it will not work. or you need to point a slice struct which contains the
			// backing array - header address, capacity and the length.
			g2.DFSRec2(num, visited, result)
		}
	}

	fmt.Println("result is : ", result, "visited : ", visited)
}

func (g2 *Graph2) DFSWithStack(start int) {

}
