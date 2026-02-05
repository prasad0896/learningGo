package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {

	dummyHead := &ListNode{-1, head}

	prevMiddle, slow, fast := dummyHead, head, head

	for fast != nil && fast.Next != nil {
		prevMiddle = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prevMiddle.Next = slow.Next
	slow = nil

	return dummyHead.Next

}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	// Declaring a linked list of 5 numbers: 1 -> 2 -> 3 -> 4 -> 5
	n5 := &ListNode{5, nil}
	n4 := &ListNode{4, n5}
	n3 := &ListNode{3, n4}
	n2 := &ListNode{2, n3}
	head := &ListNode{1, n2}

	fmt.Print("Original list: ")
	printList(head)

	newHead := deleteMiddle(head)

	fmt.Print("After deleting middle: ")
	printList(newHead)
}
