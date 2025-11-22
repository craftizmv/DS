package dp_practice

// FibRecursive - Implement the top-down approach to solve it
// use memoization
// TC = O(n)
// SC = for the recursive stack  =~ O(n)
func FibRecursive(n int, memoizedSum []int) int {
	// base cases

	if n == 0 || n == 1 {
		return n
	}

	return FibRecursive(n-1, memoizedSum) + FibRecursive(n-2, memoizedSum)
}

// FibRecursive - Implement the top-down approach to solve it
// use memoization
// TC = O(n)
// SC = O(n) for the recursive stack  and O(n) to store the results, Total =~ O(n)
func FibRecursiveMem(n int, memoizedSum []int) int {
	// base cases

	if n == 0 || n == 1 {
		return n
	}

	if memoizedSum[n] != -1 {
		return memoizedSum[n]
	}

	memoizedSum[n] = FibRecursive(n-1, memoizedSum) + FibRecursive(n-2, memoizedSum)

	return memoizedSum[n]
}

// FibUsingTabulization Bottom up approach to calculate the dp sum.
// use memoization
// TC = O(n)
// SC = O(n) for array as we are using the array(1D Table) =~ O(n)
func FibUsingTabulization(n int) int {
	// base case
	if n == 0 || n == 1 {
		return n
	}

	// init 1D table = array
	table1D := make([]int, n+1)
	table1D[0] = 0
	table1D[1] = 1

	for i := 2; i <= n; i++ {
		table1D[i] = table1D[i-1] + table1D[i-2]
	}

	// return the last value
	return table1D[n]
}

// FibTabulizationWithSpaceOptimization - Fib doing space optimisation
// don't use 1D table .. try to see if we can simply use just variables
// use memoization
// TC = O(n)
// SC = O(1) as we are just using constant variables.
func FibTabulizationWithSpaceOptimization(n int) int {
	//base condition
	if n == 0 || n == 1 {
		return n
	}

	var prev2, prev1 = 0, 1
	for i := 2; i <= n; i++ {
		sum := prev2 + prev1
		// shift vars
		prev2 = prev1
		prev1 = sum
	}

	return prev1
}
