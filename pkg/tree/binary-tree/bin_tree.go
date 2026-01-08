package binary_tree

import (
	"errors"
	"fmt"
	"math"

	"github.com/lossdev/stack"
)

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
// ######################### Traversals Recursive ########################
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
// ##################### Traversals using stack ##########################
// #######################################################################
/*
 DFSPreOrderStack - shows the approach to do pre-order traversal using a stack.
 - Approach = how to create stack.
 - print root.. push right and then left ...
 - above will act as a seed.

 - in a continous while loop ... pop print .. add to result push child ..right then left
 - find its child and push ... [no need to]
 -
*/
func (bt *BinaryTree) DFSPreOrderStack(currentNode *Node) {
	if currentNode == nil {
		return
	}

	st := stack.NewGenericStack()
	st.Push(currentNode)

	for st.Size() > 0 {
		// pop & print
		t, _ := st.Pop()
		top := t.(*Node)

		fmt.Println("Item : ", top.Data)

		// push childrens
		if top.Right != nil {
			st.Push(top.Right)
		}

		if top.Left != nil {
			st.Push(top.Left)
		}
	}
}

/*
 DFSInorderTraversal - Keep on pushing left until the left is
 null -> then pop and print the top -> check if right exists, if yes, then push it.
 -> continue the loop which should keep on pushing the left of the subtree

 output : it just prints
*/

func (bt *BinaryTree) DFSInOrderStack(root *Node) {
	if root == nil {
		return
	}

	st := stack.NewGenericStack()
	// pushing the root as a seed.
	st.Push(root)

	// jab tak stack me kuch bhara hai .. tab tak continue
	for st.Size() > 0 {
		// peek the top and check if left exists
		// t, _ := st.Peek()
		// top := t.(*Node)
		top := root
		// if left exits
		if top.Left != nil {
			st.Push(top.Left)
			root = top.Left
		} else {
			// 1. Pop and Print
			t, _ := st.Pop()
			top := t.(*Node)
			fmt.Println("Data is : ", top.Data)

			if top.Right != nil {
				st.Push(top.Right)
				root = top.Right
			}
		}
	}
}

func (bt *BinaryTree) DFSPostOrderTwoStack(root *Node) {
	if root == nil {
		return
	}

	st1 := stack.NewGenericStack()
	// pushing the root as a seed.
	st1.Push(root)

	st2 := stack.NewGenericStack()

	// jab tak stack me kuch bhara hai .. tab tak continue
	for st1.Size() > 0 {
		t, _ := st1.Pop()
		top := t.(*Node)

		// 1. Push to stack2
		st2.Push(top)

		// push left and then right of top element to stack
		if top.Left != nil {
			st1.Push(top.Left)
		}

		if top.Right != nil {
			st1.Push(top.Right)
		}
	}

	// pop st2 to print everything in post order
	for st2.Size() > 0 {
		t, _ := st2.Pop()
		fmt.Println("Node is : ", t.(*Node).Data)
	}
}

// Reference : https://youtu.be/kcTcfOWFizA?si=BKQLyUvQpKGJFavq
func (bt *BinaryTree) DFSPostOrderOneStack(root *Node) {
	// seed position
	st := stack.NewGenericStack()
	curr := root

	// while true loop
	for {
		if curr != nil {
			st.Push(curr)
			curr = curr.Left
		} else {
			if st.Size() == 0 {
				// stack is empty - no need to further process it.
				break
			}

			// below means that we need to check the right subtree of the top of the stack as either the left
			// subtree has been processed or it was found to be nil and now the pointer need to be on the right
			// of the top of the stack.
			t, _ := st.Peek()
			curr = t.(*Node).Right
			if curr == nil {
				// right is nil .. now it means that we need to pop and print what is on top of the stack.
				// maintain a last variable to track if the prev processed element is stack top right ..if yes
				// it means that right subtree has been processed and we can process the left subtree
				var last *Node
				for st.Size() > 0 {
					tempTop, _ := st.Peek()
					if tempTop.(*Node).Right == last {
						// 1. Pop -> Print -> set last to popped element
						top, _ := st.Pop()
						last = top.(*Node)
						fmt.Print("Node Data is :", last.Data, "  \n")
					} else {
						break
					}
				}
			}
		}

	}

}

// #######################################################################
// ##################### utility functions ###############################
// #######################################################################

// TC : O(N) - You would have to visit all the nodes and its branches to determine
// the complexity.
func (bt *BinaryTree) MaxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	return 1 + max(bt.MaxDepth(root.Left), bt.MaxDepth(root.Right))
}

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
