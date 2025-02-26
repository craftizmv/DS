package traversals

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


