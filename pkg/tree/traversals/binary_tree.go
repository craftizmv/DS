package traversals

import (
	"errors"
	"fmt"
)

// TODO: make it generic with golang
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

// MorrisInOrderTraversal : step 1 : make linking for inorder predecessors
// step 2 : we will print when we reach a node where it does have any Left or Right node.
// or Right of the previously found inorder predecessor is equal to root.
func (bt *BinaryTree) MorrisInOrderTraversal(root *Node) error {
	if root == nil {
		return errors.New("root is nil")
	}
	current := root

	for current != nil {

		// below condition means that .. you have processed the entire left subtree.
		if current.Left == nil {
			fmt.Println(current.Data)
			current = current.Right
		} else {
			// step 1 : find the inorder predecessor for the now's current node
			// This is done to later come back to parent. [In this way, the parent linking]
			// is established.
			currentPredecessor := current.Left
			for currentPredecessor.Right != nil && currentPredecessor.Right != current {
				currentPredecessor = currentPredecessor.Right
			}

			if currentPredecessor.Right == current {
				// Condition : of prev linking
				// step 1 : print current
				// step 2 : break chain
				// step 3 : move current to current's Right
				fmt.Println(current.Data)
				currentPredecessor.Right = nil
				current = current.Right
			} else {
				// make linking
				currentPredecessor.Right = current
				// step 2: move the current to current.Left.. this is done to navigate to the Left
				current = current.Left
			}
		}
	}

	return nil
}
