# Graph Notes

###### Types of Graph traversal
1. BFS
2. DFS

###### Mathematical Representation of Graph
G = (V,E)

#### Ways to represent a Graph :
- Adjacency List (AL)
- Adjacency Matrix. (AM)
  - Directed and Undirected graph.
  - [AM example](https://opendatastructures.org/ods-java/12_1_AdjacencyMatrix_Repres.html)
    
###### AM Properties:
1. addEdge(i,j) - ***O(1)***
2. removeEdge(i,j) - ***O(1)***
3. hasEdge(i,j) - ***O(1)***
4. inEdges(i) - ***O(n)***
5. outEdges(i) - ***O(n)***


###### AL vs  AM (Comparison)
1. AM gives benefit in finding whether a path exists between two nodes in O(1).
2. AM occupies space O(V<sup>2</sup>) and can be costly. But makes sense when the graph is dense.
3. AL is more efficient in terms of space but to search whether a path exist or not is O(n).


### Main Properties of Graph :
1. Finding path from one node to another.
2. Finding the shortest path from one node to another.
3. Find whether the path exists.
