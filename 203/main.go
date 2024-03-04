package main

// Time Complexity: O(N)
// Space Complexity: O(1)

type ListNode struct {
	Val  int
	Next *ListNode
}

// func removeElements(head *ListNode, val int) *ListNode {

// 	// 處理head是val的情況
// 	for head != nil && head.Val == val {
// 		head = head.Next
// 	}

// 	curr := head

// 	for curr != nil && curr.Next != nil {
// 		if curr.Next.Val == val {
// 			curr.Next = curr.Next.Next
// 		} else {
// 			curr = curr.Next
// 		}
// 	}

// 	return head
// }

// 使用虛擬頭節點, 簡化處理過程
func removeElements(head *ListNode, val int) *ListNode {
	dummyhead := &ListNode{}
	dummyhead.Next = head
	curr := dummyhead

	for curr != nil && curr.Next != nil {
		// 需要curr.Val 操作,確保不為nil
		// 需要curr.Next 操作,確保不為nil
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return dummyhead.Next
}
