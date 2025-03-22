package main

import (
	"fmt"
	dp_practice "github.com/craftizmv/DS/pkg/dp-practice"
)

func main() {
	//result := FibRecursive(7, make([]int, 8))
	//result := FibUsingTabulization(0)
	//result := FibTabulizationWithSpaceOptimization(7)
	//println(result)

	//numWaysClimb := CountWaysToNthStair(6, 0)
	//println(numWaysClimb)

	// Count ways with DP solution
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	minCost := dp_practice.CountTheCostToReachNthStair(cost)
	fmt.Println("Min Cost to Climb nth stair is", minCost)

	minCostWithDp := dp_practice.CountTheCostToReachNthStairWithDp(cost)
	fmt.Println("Min Cost to Climb nth stair with DP Sol is : ", minCostWithDp)

	minCostWithTabulation := dp_practice.CountTheCostToReachNthStairWithTabulation1D(cost)
	fmt.Println("Min Cost to Climb nth stair with DP Sol with Tabulation : ", minCostWithTabulation)

	minCostWithTabulationSOptim := dp_practice.CountTheCostToReachNthStairWithTabulationSpaceOptim(cost)
	fmt.Println("Min Cost to Climb nth stair with DP Sol with Space Optim Tabulation : ", minCostWithTabulationSOptim)

}
