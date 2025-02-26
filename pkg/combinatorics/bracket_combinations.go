package combinatorics

import (
	"fmt"
	"time"
)

type name struct {
}

type Vertex struct{
	Start string
	End string
	Weight int16
}

type Graph struct {
	AdjMap map[string]int
}




func BracketCombinations(num int) int {
	for i := 0; i < ; i++ {
		
	}

	graph := Graph{
		AdjMap: make(map[string]int,0),
	}

	vertexAddr := &Vertex {
		Start: "1",
		End: "2",
		Weight: 3,
	}

	graph.AdjMap["12"] = vertexAddr


	// Zero brackets gives us zero combinations
	if num == 0 {
		return 0
	}

	combinationsCount := 0

	// Functoin that will walk all the pathes with brackets
	var walker func(int, int)

	walker = func(openBrackets, clozedBrackets int) {
		time.Sleep(time.Second * 5)
		fmt.Println("OB : ", openBrackets, "CB: ", clozedBrackets)

		if openBrackets == 4 {
			fmt.Println("Check this condition")
		}

		// if we reached the end of path
		if openBrackets == num && clozedBrackets == num {
			fmt.Println("condition reached .. OB and CB are same to num : ", num)
			combinationsCount++ // scope captured variable
			return
		}

		if clozedBrackets > openBrackets {
			return // wrong path (cannot close bracket before openning)
		}

		if clozedBrackets > num {
			return // wrong path (cannot close brackets more than total `num` supply)
		}

		if openBrackets > num {
			return // wrong path (cannot open brackets more than total `num` supply)
		}

		// Path to open bracket
		walker(openBrackets+1, clozedBrackets)

		// Path to close bracket
		walker(openBrackets, clozedBrackets+1)
	}

	walker(0, 0)

	return combinationsCount
}


func giveMeShortestPath(adjMap map[string]*Vertex, start, end string) string {
	// get connected vertex
	if connections, ok := adjMap[start];!ok {

	} else {
		con
	}

}



// ===

package main
import "fmt"

type Vertex struct{
	Start string
	End string
	Weight int16
}

type Graph struct {
	AdjMap map[string]*[]Vertex
}

func WeightedPath(strArr []string) string {
	if len(strArr) <= 3 {
		// invalid data
		return "-1"
	}

	graph := Graph{
		AdjMap: make(map[string][]*Vertex,0),
	}

	// get the length of the array
	lenNodes := strArr[0]

	// we need to form the
	for i:= lenNodes+1; i < len(strArr);i++ {
		currentVertex := strArr[i]
		connectionArr := strings.Split(currentVertex,"|")

		if len(connectionArr) != 3 {
			// invalid data case
			return "-1"
		}

		vertex := &Vertex {
			Start: connectionArr[0],
			End: connectionArr[1],
			Weight: connectionArr[2],
		}

		key := vertex.Start + "-"+ vertex.End
		revKey := vertex.End + "-" + vertex.Start

		if adjList, ok := graph.AdjMap[key]; !ok {
			if adjList == nil {
				adjList = make([]*Vertex,0)
			}
			adjList = append(adjList, vertex)
		} else {
			adjList = append(adjList, vertex)
		}
		adjList[revKey] = adjList
		adjList[key] = adjList

	}

	// finding the shortest path
	giveMeShortestPath(graph.AdjMap, strArr[1], strArr[lenNodes])


	// code goes here
	return strArr[0];

}



func giveMeShortestPath(adjMap map[string]*Vertex, start, end string) string {
	// get connected vertex
	if connectionList, ok := adjMap[start];!ok {

	} else {

	}

}



func main() {

	// do not modify below here, readline is our function
	// that properly reads in the input for you
	fmt.Println(WeightedPath(readline()))

}