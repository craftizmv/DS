package traversals

import (
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

// captureOutput captures stdout produced by the provided function and returns it as a string.
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	_ = w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	_ = r.Close()

	return buf.String()
}

// parseOutputToInts parses the captured output from traversal functions and
// returns the sequence of node values in the order they were printed.
func parseOutputToInts(out string) []int64 {
	lines := strings.Split(out, "\n")
	var result []int64

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}

		last := fields[len(fields)-1]
		v, err := strconv.ParseInt(last, 10, 64)
		if err != nil {
			continue
		}
		result = append(result, v)
	}

	return result
}

// buildBalancedTree builds the sample balanced tree used in cmd/dsa-trees/main.go
// and returns its root. The structure is:
//
//           1
//        /     \
//       2       6
//      / \     / \
//     3   4   9   7
//        / \
//       8   5
func buildBalancedTree() *Node {
	root := &Node{Data: 1}

	root.Left = &Node{Data: 2}
	root.Right = &Node{Data: 6}

	root.Right.Left = &Node{Data: 9}
	root.Right.Right = &Node{Data: 7}

	root.Left.Left = &Node{Data: 3}
	root.Left.Right = &Node{Data: 4}
	root.Left.Right.Left = &Node{Data: 8}
	root.Left.Right.Right = &Node{Data: 5}

	return root
}

func TestDFSInOrderStack_BalancedTree(t *testing.T) {
	root := buildBalancedTree()
	bt := NewBinaryTree(root)

	out := captureOutput(func() {
		bt.DFSInOrderStack(root)
	})

	got := parseOutputToInts(out)
	want := []int64{3, 2, 8, 4, 5, 1, 9, 6, 7}

	if len(got) != len(want) {
		t.Fatalf("expected %d nodes, got %d; output=%q", len(want), len(got), out)
	}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("in-order mismatch at index %d: want %d, got %d (output=%q)", i, want[i], got[i], out)
		}
	}
}

func TestDFSInOrderStack_EmptyTree(t *testing.T) {
	bt := NewBinaryTree(nil)

	out := captureOutput(func() {
		bt.DFSInOrderStack(nil)
	})

	got := parseOutputToInts(out)
	if len(got) != 0 {
		t.Fatalf("expected no output for empty tree, got %v (raw=%q)", got, out)
	}
}

func TestDFSPostOrderTwoStack_BalancedTree(t *testing.T) {
	root := buildBalancedTree()
	bt := NewBinaryTree(root)

	out := captureOutput(func() {
		bt.DFSPostOrderTwoStack(root)
	})

	got := parseOutputToInts(out)
	want := []int64{3, 8, 5, 4, 2, 9, 7, 6, 1}

	if len(got) != len(want) {
		t.Fatalf("expected %d nodes, got %d; output=%q", len(want), len(got), out)
	}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("post-order mismatch at index %d: want %d, got %d (output=%q)", i, want[i], got[i], out)
		}
	}
}

func TestDFSPostOrderTwoStack_EmptyTree(t *testing.T) {
	bt := NewBinaryTree(nil)

	out := captureOutput(func() {
		bt.DFSPostOrderTwoStack(nil)
	})

	got := parseOutputToInts(out)
	if len(got) != 0 {
		t.Fatalf("expected no output for empty tree, got %v (raw=%q)", got, out)
	}
}

func TestDFSInOrderStack_SkewedTree(t *testing.T) {
	// Right-skewed tree: 1 -> 2 -> 3
	root := &Node{Data: 1}
	root.Right = &Node{Data: 2}
	root.Right.Right = &Node{Data: 3}

	bt := NewBinaryTree(root)

	out := captureOutput(func() {
		bt.DFSInOrderStack(root)
	})

	got := parseOutputToInts(out)
	want := []int64{1, 2, 3}

	if len(got) != len(want) {
		t.Fatalf("expected %d nodes, got %d; output=%q", len(want), len(got), out)
	}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("in-order mismatch for skewed tree at index %d: want %d, got %d (output=%q)", i, want[i], got[i], out)
		}
	}
}
