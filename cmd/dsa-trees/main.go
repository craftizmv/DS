package main

import (
	"fmt"

	binary_tree "github.com/craftizmv/DS/pkg/tree/binary-tree"
)

// morris traversal
func main() {
	// root := &binary_tree.Node{
	// 	Data: 1,
	// }
	// root.Left = &binary_tree.Node{
	// 	Data: 2,
	// }
	// root.Right = &binary_tree.Node{
	// 	Data: 6,
	// }
	// root.Right.Left = &binary_tree.Node{
	// 	Data: 9,
	// }
	// root.Right.Right = &binary_tree.Node{
	// 	Data: 7,
	// }

	// root.Left.Left = &binary_tree.Node{
	// 	Data: 3,
	// }
	// root.Left.Right = &binary_tree.Node{
	// 	Data: 4,
	// }
	// root.Left.Right.Left = &binary_tree.Node{
	// 	Data: 8,
	// }
	// root.Left.Right.Right = &binary_tree.Node{
	// 	Data: 5,
	// }

	// bt := binary_tree.NewBinaryTree(root)
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

	// bt.DFSPostOrderTwoStack(root)

	// bt.DFSPostOrderOneStack(root)

	// fmt.Println("Depth is ", bt.MaxDepth(root))

	skewTree, root := sampleSkewTree()
	fmt.Println("Is Balanced ", skewTree.IsBalancedTree(root))

}

func sampleSkewTree() (*binary_tree.BinaryTree, *binary_tree.Node) {

	root := &binary_tree.Node{
		Data: 1,
	}
	root.Left = &binary_tree.Node{
		Data: 2,
	}
	root.Left.Left = &binary_tree.Node{
		Data: 6,
	}
	root.Left.Left.Left = &binary_tree.Node{
		Data: 9,
	}
	root.Left.Left.Left.Right = &binary_tree.Node{
		Data: 7,
	}

	root.Left.Left = &binary_tree.Node{
		Data: 3,
	}
	root.Left.Right = &binary_tree.Node{
		Data: 4,
	}
	root.Left.Right.Left = &binary_tree.Node{
		Data: 8,
	}
	root.Left.Right.Right = &binary_tree.Node{
		Data: 5,
	}

	bt := binary_tree.NewBinaryTree(root)
	return bt, root
}
