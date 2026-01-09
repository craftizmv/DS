package medium

// Node : TODO - make it generic with golang
type Node struct {
	Data  int64
	Left  *Node
	Right *Node
}

type BinaryTree struct {
	root *Node
}

func NewBinaryTree(root *Node) *BinaryTree {
	return &BinaryTree{
		root: root,
	}
}

// TC : O(N) - You would have to visit all the nodes and its branches to determine
// the complexity.
func (bt *BinaryTree) MaxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	return 1 + max(bt.MaxDepth(root.Left), bt.MaxDepth(root.Right))
}
