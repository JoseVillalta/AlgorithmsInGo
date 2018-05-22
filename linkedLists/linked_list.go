package linkedlists

import "fmt"

// ListNode contains the element of the linked list.
type ListNode struct {
	data int
	next *ListNode
}

// LinkedList class to learn go.
type LinkedList struct {
	head *ListNode
	back *ListNode
}

// Inserts a value at the end of the list
func (l *LinkedList) Insert(val int) {
	node := new(ListNode)
	node.data = val

	if l.head == nil {
		l.head = node
	}
	if l.back == nil {
		l.back = node
	} else {
		l.back.next = node
		l.back = node
	}

}

func (l *LinkedList) Remove(val int) {

}

func (l * LinkedList) PrintAllValues() {
	nodePtr := l.head
	for nodePtr != nil {
		fmt.Println(nodePtr.data)
		nodePtr = nodePtr.next
	}
}
