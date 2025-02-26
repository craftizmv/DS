package traversals

import (
	"errors"
	"fmt"
)

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

// BFS - This is solved using the queue.
func (bt *BinaryTree) BFS() {

}

// #######################################################################
// ############################ DFS RECURSIVE ############################
// #######################################################################

// DFSPreOrderRecursive (DLR) - This can be implemented using stack or using recursion
func (bt *BinaryTree) DFSPreOrderRecursive(currentNode *Node) {
	// base condition
	if currentNode != nil {
		// visit node
		println(currentNode.Data)
	} else {
		return
	}

	// base condition
	if currentNode.Left == nil && currentNode.Right == nil {
		// reached leaf node.
		return
	}

	// visit left node recursively
	bt.DFSPreOrderRecursive(currentNode.Left)

	// visit right node recursively

	bt.DFSPreOrderRecursive(currentNode.Right)
}

// DFSInOrderRecursive (LDR) - Implemented using recursion.
func (bt *BinaryTree) DFSInOrderRecursive(currentNode *Node) {
	if currentNode == nil {
		return
	}

	if currentNode.Left != nil {
		// call recursively
		bt.DFSInOrderRecursive(currentNode.Left)
	}

	// print the current Node
	println(currentNode.Data)

	if currentNode.Right != nil {
		bt.DFSInOrderRecursive(currentNode.Right)
	}
}

// DFSPostOrderRecursive (LRD) - Implemented using recursion.
func (bt *BinaryTree) DFSPostOrderRecursive(currentNode *Node) {
	if currentNode == nil {
		return
	}

	if currentNode.Left != nil {
		bt.DFSPostOrderRecursive(currentNode.Left)
	}

	if currentNode.Right != nil {
		bt.DFSPostOrderRecursive(currentNode.Right)
	}

	println(currentNode.Data)
}

// #######################################################################
// ############################ DFS USING STACK ############################
// #######################################################################
/*
 - Approach = how to create stack.
 -
*/
