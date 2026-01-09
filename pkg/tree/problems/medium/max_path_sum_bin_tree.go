package medium

// MaxPathValueSum - returns the maximum path sum of a binary tree.
// Problem - inspired by Height of the Tree and Diameter of a tree.
func (bt *BinaryTree) MaxPathValueSum(root *Node, maxi *int) int {
	if root == nil {
		return 0
	}

	lSum := bt.MaxPathValueSum(root.Left, maxi)
	rSum := bt.MaxPathValueSum(root.Right, maxi)

	*maxi = max(*maxi, lSum+rSum+int(root.Data))

	return int(root.Data) + max(lSum, rSum)
}
