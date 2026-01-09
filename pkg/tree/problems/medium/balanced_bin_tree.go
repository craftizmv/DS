package medium

import (
	"math"
)

// IsBalancedTree - Checks if the tree is balanced or not.
// TC : O(N)
func (bt *BinaryTree) IsBalancedTree(root *Node) bool {
	return bt.isBalanced(root) != -1
}

func (bt *BinaryTree) isBalanced(root *Node) int {
	if root == nil {
		return 0
	}

	leftHeight := 1 + bt.isBalanced(root.Left)
	rightHeight := 1 + bt.isBalanced(root.Right)

	if leftHeight == -1 || rightHeight == -1 {
		// it means that either the left or the right subtree
		// is not balanced.
		return -1
	}

	// if abs difference is greater than 1 then return -1
	if math.Abs(float64(rightHeight-leftHeight)) > 1 {
		return -1
	}

	return max(leftHeight, rightHeight)
}
