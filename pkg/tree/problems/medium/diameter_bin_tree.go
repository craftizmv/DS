package medium

// No need to associate it as a method of the binary tree we can disassociate also ..

func (bt *BinaryTree) Diameter(root *Node) int {
	if root == nil {
		return 0
	}

	result := 0
	diameter(root, &result)

	return result
}

// This code replicates getting the max-height of the bin tree
// just in the process we try to determine the max curvature via
// a point which is called the diameter.
// return the max height of the bin tree.
// TC : O(n)
func diameter(root *Node, result *int) int {
	if root == nil {
		return 0
	}

	lh := diameter(root.Left, result)
	rh := diameter(root.Right, result)

	*result = max(*result, lh+rh)

	return 1 + max(lh, rh)
}

// TC : O(n²)
func (bt *BinaryTree) Diameter2(root *Node, maxi *int) int {
	if root == nil {
		return 0
	}

	// this implementation going to the left and right together takes
	// O(n) TC for every recursion call which itself causes O(n)
	lh := bt.MaxDepth(root.Left)
	rh := bt.MaxDepth(root.Right)

	*maxi = max(*maxi, lh+rh)

	bt.Diameter2(root.Left, maxi)
	bt.Diameter2(root.Right, maxi)

	return *maxi
}
