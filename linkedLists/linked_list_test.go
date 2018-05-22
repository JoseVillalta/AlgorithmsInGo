package linkedlists

import (
	"testing"
	"fmt"
)

func TestLinkedList_Insert(t *testing.T) {
	myList := new(LinkedList)
	myList.Insert(1)
	fmt.Println("Inserted a value")
	myList.Insert(2)
	myList.Insert(3)
	myList.PrintAllValues()

}
