package main

import "fmt"

func main() {

}

type ListNode struct {
	Val  int       // Data
	Next *ListNode // Pointer to next node
}

func printList(head *ListNode) {
	current := head

	for current != nil {
		fmt.Printf("%d → ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func insertAtBeginning(head *ListNode, val int) *ListNode {
	newNode := &ListNode{Val: val}
	newNode.Next = head
	return newNode // New head
}

func insertAtEnd(head *ListNode, val int) *ListNode {
	newNode := &ListNode{Val: val}

	// Edge case: empty list
	if head == nil {
		return newNode
	}

	// Traverse to last node
	current := head
	for current.Next != nil {
		current = current.Next
	}

	// Attach new node
	current.Next = newNode
	return head
}

func insertAtPosition(head *ListNode, val int, pos int) *ListNode {
	newNode := &ListNode{Val: val}

	// Insert at beginning
	if pos == 0 {
		newNode.Next = head
		return newNode
	}

	// Traverse to position-1
	current := head
	for i := 0; i < pos-1 && current != nil; i++ {
		current = current.Next
	}

	// If position is valid
	if current != nil {
		newNode.Next = current.Next
		current.Next = newNode
	}

	return head
}

func deleteAtBeginning(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	return head.Next // New head
}

func deleteAtEnd(head *ListNode) *ListNode {
	// Empty list
	if head == nil {
		return nil
	}

	// Single node
	if head.Next == nil {
		return nil
	}

	// Traverse to second-to-last node
	current := head
	for current.Next.Next != nil {
		current = current.Next
	}

	// Remove last node
	current.Next = nil
	return head
}

func deleteByValue(head *ListNode, val int) *ListNode {
	// Empty list
	if head == nil {
		return nil
	}

	// Delete head if it matches
	if head.Val == val {
		return head.Next
	}

	// Find node before the one to delete
	current := head
	for current.Next != nil && current.Next.Val != val {
		current = current.Next
	}

	// If found, skip the node
	if current.Next != nil {
		current.Next = current.Next.Next
	}

	return head
}
