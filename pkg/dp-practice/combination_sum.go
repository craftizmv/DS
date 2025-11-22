package dp_practice

import "sort"

var result [][]int

// Given - an array containing elements
// find - the number of unique combinations where the
// sum of elements in combination is equal to the given target.
// Conditions:
// 1. The same number may be choosen any number of times to form the target.
// 2. All the numbers will the positive integers.
// 3. Elements in the comb should be in the non-decreasing order.
// 4. Comb themselves must be sorted in ascending order.
func FindAllUniqueCombinationsForTarget(input []int, target int) [][]int {

	// sort
	sort.Ints(input)

	// remove duplicate.
	input = UniqueInts(input)

	tempPath := make([]int, 0)
	// rSolve(input, target, tempPath)
	rSolve2(input, target, tempPath, 0)
	return result
}

///////////////////////////////////////////////////////
////////////////////  Helper Funcs ////////////////////
///////////////////////////////////////////////////////

func UniqueInts(nums []int) []int {
	seen := make(map[int]bool)
	out := make([]int, 0, len(nums))
	for _, n := range nums {
		if !seen[n] {
			seen[n] = true
			out = append(out, n)
		}
	}
	return out
}

///////////////////////////////////////////////////////
////////////////////  Recur Solve Funcs ////////////////////
///////////////////////////////////////////////////////

// rSolve - iterates over all the elements
// as a branch to explore the path
// Problems in the code:
// Major : path as a reference is manipulated after addition
// into the slides.
func rSolve(input []int, target int, path []int) {
	// base condition - we need to add the path to the final answer.
	if target == 0 {
		result = append(result, path)
	}

	if target < 0 {
		// we need to just ignore the path
		return
	}

	// we will explore all the branches for the given target.
	for i := range input {
		v := input[i]

		// TODO :  Need to work on understanding the boundry conditions.
		if target-v >= 0 {
			// 1. append in the path,
			path = append(path, v)
			// 2. recur call with new target and path.
			rSolve(input, target-v, path)
			// 3. remove this element from the path,
			// so that other option can be called.(backtrack)
			path = path[:len(path)-1]
		} else {
			break // check if we can break.
			// as the input is sorted so other solutions
			// will also result in break.
		}
	}
}

func rSolve2(input []int, target int, path []int, idx int) {
	if target == 0 {
		result = append(result, path)
	}

	if target < 0 {
		// not a valid path.
		return
	}

	// from the current index we want to see all the combinations.
	for j := idx; j < len(input); j++ {
		v := input[j]

		if target-v >= 0 {
			// 1. append in the path,
			path = append(path, v)
			// 2. recur call with new target and path.
			rSolve2(input, target-v, path, j)
			// 3. remove this element from the path,
			// so that other option can be called.(backtrack)
			path = path[:len(path)-1]
		} else {
			break
		}
	}
}

// GPT : polished code
// FindAllUniqueCombinationsForTarget finds all unique combinations
// where the sum equals the given target. Same number can be reused.
func FindAllUniqueCombinationsForTarget2(input []int, target int) [][]int {
	sort.Ints(input)
	input = UniqueInts(input)

	var result [][]int
	var path []int

	var dfs func(start, remain int)
	dfs = func(start, remain int) {
		if remain == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}

		for i := start; i < len(input); i++ {
			v := input[i]
			if v > remain {
				break // sorted → no need to check further
			}
			path = append(path, v)
			dfs(i, remain-v)          // stay on i → reuse allowed
			path = path[:len(path)-1] // backtrack
		}
	}

	dfs(0, target)
	return result
}
