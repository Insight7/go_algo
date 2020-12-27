package linkedlist

import (
	"fmt"
)

// DoublyLinkedList definition
type DoublyLinkedList struct {
	Head   *DoublyNode
	length int
}

// DoublyNode is node for DoublyLinkedList
type DoublyNode struct {
	Data int
	Next *DoublyNode
	Prev *DoublyNode
}

// Basic operations

// Length get the length of the DoublyLinkedList
func (linkedList *DoublyLinkedList) Length() int {
	if linkedList == nil {
		return 0
	}
	return linkedList.length
}

// IsEmpty returns if the DoublyLinkedList is empty
func (linkedList *DoublyLinkedList) IsEmpty() bool {
	return linkedList.Length() == 0
}

// InsertAtStart is used to insert element at start of DoublyLinkedList
func (linkedList *DoublyLinkedList) InsertAtStart(data int) {
	node := DoublyNode{
		Data: data,
		Next: linkedList.Head,
	}
	if linkedList.Head != nil {
		linkedList.Head.Prev = &node
	}
	linkedList.Head = &node
	linkedList.length++
}

// InsertAtEnd is used to insert element at end of DoublyLinkedList
func (linkedList *DoublyLinkedList) InsertAtEnd(data int) {
	node := DoublyNode{
		Data: data,
	}
	if linkedList.IsEmpty() {
		linkedList.InsertAtStart(data)
		return
	}
	head := linkedList.Head
	for head.Next != nil {
		head = head.Next
	}
	node.Prev = head
	head.Next = &node
	linkedList.length++
}

// InsertAt is used to insert element at index position of DoublyLinkedList
func (linkedList *DoublyLinkedList) InsertAt(index uint, data int) {
	if linkedList.IsEmpty() {
		linkedList.InsertAtStart(data)
		return
	}
	if linkedList.length <= int(index)+1 {
		linkedList.InsertAtEnd(data)
		return
	}

	temp := linkedList.Head
	for temp.Next != nil && index > 1 {
		index--
		temp = temp.Next
	}
	node := &DoublyNode{
		Data: data,
		Next: temp.Next,
		Prev: temp,
	}
	temp.Next.Prev = node
	temp.Next = node
	linkedList.length++
}

// DeleteAtStart is used to delete element at start position of DoublyLinkedList
func (linkedList *DoublyLinkedList) DeleteAtStart() {
	if linkedList.IsEmpty() {
		panic("Linked list is empty hence cannot delete")
	}
	linkedList.length--
	if linkedList.length == 0 {
		linkedList.Head = nil
		return
	}
	linkedList.Head.Next.Prev = nil
	linkedList.Head = linkedList.Head.Next
}

// DeleteAtEnd is used to delete element at end position of DoublyLinkedList
func (linkedList *DoublyLinkedList) DeleteAtEnd() {
	if linkedList.IsEmpty() {
		panic("Linked list is empty hence cannot delete")
	}
	if linkedList.Length() == 1 {
		linkedList.DeleteAtStart()
		return
	}
	head := linkedList.Head
	for head.Next.Next != nil {
		head = head.Next
	}
	head.Next = nil
	linkedList.length--

}

// DeleteAt is used to delete element at index position of DoublyLinkedList
func (linkedList *DoublyLinkedList) DeleteAt(index uint) {
	if linkedList.IsEmpty() {
		panic("Linked list is empty hence cannot delete")
	}
	if index == 0 {
		linkedList.DeleteAtStart()
		return
	} else if index >= uint(linkedList.Length()-1) {
		linkedList.DeleteAtEnd()
		return
	}

	head := linkedList.Head
	temp := head
	for temp.Next != nil && index > 0 {
		index--
		temp = temp.Next
	}
	temp.Next.Prev = temp.Prev
	temp.Prev.Next = temp.Next
	linkedList.length--
}

// IndexOf is used to get index[0..n] of a data in the DoublyLinkedList
func (linkedList *DoublyLinkedList) IndexOf(data int) int {
	index := 0
	current := linkedList.Head
	for true {
		if current.Data == data {
			return index
		}
		index++
		if current.Next == nil {
			// data not found in the linkedList
			return -1
		}
		current = current.Next
	}
	return -1
}

// Contains is used to check if a data is in the DoublyLinkedList
func (linkedList *DoublyLinkedList) Contains(data int) bool {
	return linkedList.IndexOf(data) >= 0
}

func (node DoublyNode) String() string {
	current := node
	preetyString := "Nodes{ "
	for true {
		// check for last node
		if current.Data == 0 && current.Next == nil {
			preetyString += " nil }"
			return preetyString
		}
		// edge case for length 1
		if current.Next == nil {
			preetyString += fmt.Sprintf("[%d] <==> nil }", current.Data)
			return preetyString
		}
		preetyString += fmt.Sprintf("[%d] <==> ", current.Data)
		current = *current.Next
	}
	return ""
}

func (linkedList DoublyLinkedList) String() string {
	return fmt.Sprintf("DoublyLinkedList{ lenght: %d, head: %s }", linkedList.length, linkedList.Head)
}
