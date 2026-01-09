package medium

func IsIdentical(t1 *Node, t2 *Node) bool {
	// indicates presence of the leaf node.
	if t1 == nil || t2 == nil {
		return t1 == t2
	}

	return (t1.Data == t2.Data) && IsIdentical(t1.Left, t2.Left) && IsIdentical(t1.Right, t2.Right)

}
