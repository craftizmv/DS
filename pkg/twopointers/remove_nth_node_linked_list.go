package twopointers

import (
	"fmt"
)

// LinkedList - container for linked list
type LinkedList struct {
	head *LinkedListNode
}

// InsertNodeAtHead - inserts node at the head of the linked list
func (l *LinkedList) InsertNodeAtHead(node *LinkedListNode) {
	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}
}

func (l *LinkedList) CreateLinkedList(lst []int) {
	for i := len(lst) - 1; i >= 0; i-- {
		node := InitLinkedListNode(lst[i])
		l.InsertNodeAtHead(node)
	}
}

// DisplayLinkedList - displays element of linked list
func (l *LinkedList) DisplayLinkedList() {
	temp := l.head
	fmt.Print("{")
	for temp != nil {
		fmt.Print(temp.data)
		temp = temp.next
		if temp != nil {
			fmt.Print(", ")
		}
	}
	fmt.Print("}")
}

// ----// ----// ----// ----// ----// ----// ----// ----// ----// ----// ----

type LinkedListNode struct {
	data int
	next *LinkedListNode
}

func NewLinkedListNode(data int, next *LinkedListNode) *LinkedListNode {
	node := new(LinkedListNode)
	node.data = data
	node.next = next
	return node
}

func InitLinkedListNode(data int) *LinkedListNode {
	node := new(LinkedListNode)
	node.data = data
	node.next = nil
	return node
}

// ----// ----// ----// ----// ----// ----// ----// ----// ----// ----// ----

// RemoveNodeNthPos - Remove Nth node from the end of a Singly Linked List
//   - Approach - two pointers
func (l *LinkedList) RemoveNodeNthPos(n int) *LinkedList {
	// Point two pointers, right and left, at head.
	right := l.head
	left := l.head

	// Move right pointer n elements away from the left pointer.
	for i := 0; i < n; i++ {
		right = right.next
		// Removal of the head node.
		if right == nil {
			l.head = l.head.next
			return l
		}
	}

	// Move both pointers until right pointer reaches the last node.
	for right.next != nil {
		right = right.next
		left = left.next
	}

	// At this point, left pointer points to (n-1)th element.
	// So link it to next to next element of left.
	left.next = left.next.next

	return l
}

func removeNthLastNode(head *LinkedListNode, n int) *LinkedListNode {

	// Point two pointers, right and left, at head.
	right := head
	left := head

	// Move right pointer n elements away from the left pointer.
	for i := 0; i < n; i++ {
		right = right.next
		// Removal of the head node.
		if right == nil {
			return head.next
		}
	}

	// Move both pointers until right pointer reaches the last node.
	for right.next != nil {
		right = right.next
		left = left.next
	}

	// At this point, left pointer points to (n-1)th element.
	// So link it to next to next element of left.
	left.next = left.next.next

	return head
}
