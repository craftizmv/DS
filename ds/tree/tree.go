package tree

// Reference : https://www.golangprograms.com/golang-program-to-implement-binary-tree.html

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

type BinaryTree struct {
	root *BinaryNode
}

// Inserts a binary tree
func (t *BinaryTree) insert(data int64) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{
			left:  nil,
			right: nil,
			data:  data,
		}
	} else {

	}

	return t
}

func (n *BinaryNode) insert(data int64) {
	if n == nil {
		return
	} else if data < n.data {
		if n.left == nil {
			n.left = &BinaryNode{
				data: data,
			}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{
				data: data,
			}
		} else {
			n.right.insert(data)
		}
	}
}
