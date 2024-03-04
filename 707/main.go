package main

// 單向
type SingleNode struct {
	Val  int
	Next *SingleNode
}

type MyLinkedList struct {
	dummyHead *SingleNode // 虛擬節點
	Size      int
}

func Constructor() MyLinkedList {
	newNode := &SingleNode{
		0,
		nil,
	}

	return MyLinkedList{
		dummyHead: newNode,
		Size:      0,
	}
}

// Time Complexity: O(1)
func (this *MyLinkedList) Get(index int) int {
	if this == nil || index < 0 || index >= this.Size {
		return -1
	}

	cur := this.dummyHead.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	return cur.Val
}

// Time Complexity: O(1)
func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &SingleNode{Val: val}
	newNode.Next = this.dummyHead.Next
	this.dummyHead.Next = newNode
	this.Size++
}

// Time Complexity: O(N)
func (this *MyLinkedList) AddAtTail(val int) {
	cur := this.dummyHead // 指向index前一個
	for cur.Next != nil {
		cur = cur.Next
	}

	newNode := &SingleNode{Val: val}
	cur.Next = newNode
	this.Size++
}

// Time Complexity: O(N)
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Size {
		return
	}

	newNode := &SingleNode{Val: val}
	cur := this.dummyHead // 指向index前一個
	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	newNode.Next = cur.Next
	cur.Next = newNode
	this.Size++
}

// Time Complexity: O(N)
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.Size {
		return
	}

	cur := this.dummyHead // 指向index前一個
	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	if cur.Next != nil {
		cur.Next = cur.Next.Next
		this.Size--
	}
}
