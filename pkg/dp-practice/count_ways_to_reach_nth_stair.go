package dp_practice

// CountWaysToNthStair - Recursive impl.
// f(k) = f(k+1) + f(k+2) where k=[0 .. n]
// Ref : https://leetcode.com/problems/climbing-stairs/
// Ref : https://www.youtube.com/watch?v=S31W3kohFDk&list=PLDzeHZWIZsTomOPnCiU3J95WufjE36wsb&index=2

func CountWaysToNthStair(nStairs int, currentStair int) int {
	// base case 1
	// yahan pe 1 return karna hai .. to go from current stair to the same stair
	if currentStair == nStairs {
		return 1
	}

	// base case 2 - invalid
	if currentStair > nStairs {
		return 0
	}

	return CountWaysToNthStair(nStairs, currentStair+1) + CountWaysToNthStair(nStairs, currentStair+2)
}
