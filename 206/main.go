package main

// Time Complixity: O(N)
// Space Complixity: O(1)

type ListNode struct {
	Val  int
	Next *ListNode
}

// func reverseList(head *ListNode) (prev *ListNode) {
// 	for head != nil {
// 		head.Next, prev, head = prev, head, head.Next
// 	}
// 	return
// }

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}
