package main

import (
	"fmt"
	"strings"
)

type Node struct {
	val   string
	left  *Node
	right *Node
}

func (node *Node) serialize() string {
	if node == nil {
		return "#"
	}

	leftSide := node.left.serialize()
	rightSide := node.right.serialize()

	ret := fmt.Sprintf("%s->%s->%s", node.val, leftSide, rightSide)
	fmt.Println(ret)
	return ret
}

func deserialize(str string) Node {
	node, _ := deserailizeNodes(strings.Split(str, "->"))
	return node
}

func deserailizeNodes(nodes []string) (Node, []string) {

	if len(nodes) == 0 {
		return Node{"", nil, nil}, nil
	}

	nextNode := nodes[0]

	if nextNode == "#" {
		return Node{"", nil, nil}, nodes[1:]
	}

	left, rem := deserailizeNodes(nodes[1:])
	right, _ := deserailizeNodes(rem)

	if left.val == "" && right.val == "" {
		return Node{nextNode, nil, nil}, nil
	}
	if left.val == "" {
		return Node{nextNode, nil, &right}, nil
	}
	if right.val == "" {
		return Node{nextNode, &left, nil}, nil
	}

	node := Node{val: nextNode, left: &left, right: &right}
	return node, nil

}

func main() {
	node := Node{"root", &Node{"left", &Node{"left.left", nil, nil}, nil}, &Node{"right", nil, nil}}
	fmt.Println(node.serialize())
	// root->left->left.left->#->#->#->right->#->#
	n := deserialize(node.serialize())
	fmt.Println(n.left.left.val == "left.left") // true
}
