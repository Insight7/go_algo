package linkedlist

import (
	"fmt"
)

// SinglyLinkedList definition
type SinglyLinkedList struct {
	Head   *SinglyNode
	length int
}

// SinglyNode is node for SinglyLinkedList
type SinglyNode struct {
	Data int
	Next *SinglyNode
}

// Basic operations

// Length get the length of the SinglyLinkedList
func (linkedList *SinglyLinkedList) Length() int {
	if linkedList == nil {
		return 0
	}
	return linkedList.length
}

// IsEmpty returns if the SinglyLinkedList is empty
func (linkedList *SinglyLinkedList) IsEmpty() bool {
	return linkedList.Length() == 0
}

// InsertAtStart is used to insert element at start of SinglyLinkedList
func (linkedList *SinglyLinkedList) InsertAtStart(data int) {
	node := SinglyNode{
		Data: data,
		Next: linkedList.Head,
	}
	linkedList.length++
	linkedList.Head = &node
}

// InsertAtEnd is used to insert element at end of SinglyLinkedList
func (linkedList *SinglyLinkedList) InsertAtEnd(data int) {
	node := SinglyNode{
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
	head.Next = &node
	linkedList.length++
}

// InsertAt is used to insert element at index position of SinglyLinkedList
func (linkedList *SinglyLinkedList) InsertAt(index uint, data int) {
	if linkedList.IsEmpty() {
		linkedList.InsertAtStart(data)
		return
	}
	head := linkedList.Head
	previous := head
	temp := head
	for temp != nil && index > 0 {
		index--
		previous = temp
		temp = temp.Next
	}
	previous.Next = &SinglyNode{
		Data: data,
		Next: temp,
	}
	linkedList.length++
}

// DeleteAtStart is used to delete element at start position of SinglyLinkedList
func (linkedList *SinglyLinkedList) DeleteAtStart() {
	if linkedList.IsEmpty() {
		panic("Linked list is empty hence cannot delete")
	}
	linkedList.length--
	linkedList.Head = linkedList.Head.Next
}

// DeleteAtEnd is used to delete element at end position of SinglyLinkedList
func (linkedList *SinglyLinkedList) DeleteAtEnd() {
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

// DeleteAt is used to delete element at index position of SinglyLinkedList
func (linkedList *SinglyLinkedList) DeleteAt(index uint) {
	if linkedList.IsEmpty() {
		panic("Linked list is empty hence cannot delete")
	}
	if index == 0 {
		linkedList.DeleteAtStart()
		return
	} else if index == uint(linkedList.Length()-1) {
		linkedList.DeleteAtEnd()
		return
	}

	head := linkedList.Head
	previous := head
	temp := head
	for temp != nil && index > 0 {
		index--
		previous = temp
		temp = temp.Next
	}
	previous.Next = previous.Next.Next
	linkedList.length--
}

// IndexOf is used to get index[0..n] of a data in the SinglyLinkedList
func (linkedList *SinglyLinkedList) IndexOf(data int) int {
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

// Contains is used to check if a data is in the SinglyLinkedList
func (linkedList *SinglyLinkedList) Contains(data int) bool {
	return linkedList.IndexOf(data) >= 0
}

func (node SinglyNode) String() string {
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
			preetyString += fmt.Sprintf("[%d] --> nil }", current.Data)
			return preetyString
		}
		preetyString += fmt.Sprintf("[%d] --> ", current.Data)
		current = *current.Next
	}
	return ""
}

func (linkedList SinglyLinkedList) String() string {
	return fmt.Sprintf("SinglyLinkedList{ lenght: %d, head: %s }", linkedList.length, linkedList.Head)
}
