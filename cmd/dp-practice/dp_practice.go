package main

import . "github.com/craftizmv/DS/pkg/dp-practice"

func main() {
	//result := FibRecursive(7, make([]int, 8))
	//result := FibUsingTabulization(0)
	//result := FibTabulizationWithSpaceOptimization(7)
	//println(result)

	numWaysClimb := CountWaysToNthStair(6, 0)
	println(numWaysClimb)
}
