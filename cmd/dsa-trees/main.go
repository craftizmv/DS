package main

import (
	"github.com/craftizmv/DS/pkg/tree/traversals"
)

// morris traversal
func main() {
	root := &traversals.Node{
		Data: 1,
	}
	root.Left = &traversals.Node{
		Data: 2,
	}
	root.Right = &traversals.Node{
		Data: 6,
	}
	root.Right.Left = &traversals.Node{
		Data: 9,
	}
	root.Right.Right = &traversals.Node{
		Data: 7,
	}

	root.Left.Left = &traversals.Node{
		Data: 3,
	}
	root.Left.Right = &traversals.Node{
		Data: 4,
	}
	root.Left.Right.Left = &traversals.Node{
		Data: 8,
	}
	root.Left.Right.Right = &traversals.Node{
		Data: 5,
	}

	bt := traversals.NewBinaryTree(root)
	//err := bt.MorrisInOrderTraversal(root)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//bt.DFSPreOrderRecursive(root)

	//bt.DFSInOrderRecursive(root)

	// bt.DFSPostOrderRecursive(root)

	// bt.DFSPreOrderStack(root)

	// bt.DFSInOrderStack(root)

	bt.DFSPostOrderTwoStack(root)

}
