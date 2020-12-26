package linkedlist

import (
	"testing"
)

func TestLenght(t *testing.T) {
	linkedlist := SinglyLinkedList{
		Head: &SinglyNode{
			Data: 1,
		},
		length: 1,
	}
	l := linkedlist.Length()
	if l != 1 {
		t.Errorf("Length did not matched. linkedlist: %s\n expected: 1, got: %d ", linkedlist, l)
	}
	linkedlist.Head.Next = &SinglyNode{2, nil}
	linkedlist.length++
	l = linkedlist.Length()
	if l != 2 {
		t.Errorf("Length did not matched. linkedlist: %s\n expected: 2, got: %d ", linkedlist, l)
	}
	linkedlist.Head.Next.Next = &SinglyNode{2, nil}
	linkedlist.length++
	l = linkedlist.Length()
	if l != 3 {
		t.Errorf("Length did not matched. linkedlist: %s\n expected: 3, got: %d ", linkedlist, l)
	}
}

func TestInsertAtStart(t *testing.T) {
	var linkedlist SinglyLinkedList
	linkedlist.InsertAtStart(1)
	l := linkedlist.Length()
	if l != 1 {
		t.Errorf("Length did not matched. linkedlist: %s\n expected: 1, got: %d ", linkedlist, l)
	}
	if linkedlist.Head.Data != 1 {
		t.Errorf("linkedlist data did not matched. linkedlist: %s\n expected: 1, got: %d ", linkedlist, linkedlist.Head.Data)
	}
	linkedlist.InsertAtStart(2)
	l = linkedlist.Length()
	if l != 2 {
		t.Errorf("Length did not matched. linkedlist: %s\n expected: 2, got: %d ", linkedlist, l)
	}
	if linkedlist.Head.Data != 2 {
		t.Errorf("linkedlist data did not matched. linkedlist: %s\n expected: 2, got: %d ", linkedlist, linkedlist.Head.Data)
	}
	linkedlist.InsertAtStart(3)
	l = linkedlist.Length()
	if l != 3 {
		t.Errorf("Length did not matched. linkedlist: %s\n expected: 3, got: %d ", linkedlist, l)
	}
	if linkedlist.Head.Data != 3 {
		t.Errorf("linkedlist data did not matched. linkedlist: %s\n expected: 3, got: %d ", linkedlist, linkedlist.Head.Data)
	}
}

func TestInsertAtEnd(t *testing.T) {
	var linkedlist SinglyLinkedList
	linkedlist.InsertAtEnd(1)
	l := linkedlist.Length()
	if l != 1 {
		t.Fatalf("Length did not matched. linkedlist: %s\n expected: 1, got: %d ", linkedlist, l)
	}
	if linkedlist.Head.Data != 1 {
		t.Errorf("linkedlist data did not matched. linkedlist: %s\n expected: 1, got: %d ", linkedlist, linkedlist.Head.Data)
	}
	linkedlist.InsertAtEnd(2)
	l = linkedlist.Length()
	if l != 2 {
		t.Fatalf("Length did not matched. linkedlist: %s\n expected: 2, got: %d ", linkedlist, l)
	}
	if linkedlist.Head.Next.Data != 2 {
		t.Errorf("linkedlist data did not matched. linkedlist: %s\n expected: 2, got: %d ", linkedlist, linkedlist.Head.Next.Data)
	}
	linkedlist.InsertAtEnd(3)
	l = linkedlist.Length()
	if l != 3 {
		t.Fatalf("Length did not matched. linkedlist: %s\n expected: 3, got: %d ", linkedlist, l)
	}
	if linkedlist.Head.Next.Next.Data != 3 {
		t.Errorf("linkedlist data did not matched. linkedlist: %s\n expected: 3, got: %d ", linkedlist, linkedlist.Head.Next.Next.Data)
	}
}

func TestInsertAtIndex(t *testing.T) {
	var linkedlist SinglyLinkedList
	// testing insert at start
	linkedlist.InsertAt(0, 1)
	l := linkedlist.Length()
	if l != 1 {
		t.Fatalf("Length did not matched. linkedlist: %s\n expected: 1, got: %d ", linkedlist, l)
	}
	if linkedlist.Head.Data != 1 {
		t.Errorf("linkedlist data did not matched. linkedlist: %s\n expected: 1, got: %d ", linkedlist, linkedlist.Head.Data)
	}
	//testing insert at end
	linkedlist.InsertAt(1, 3)

	l = linkedlist.Length()
	if l != 2 {
		t.Fatalf("Length did not matched. linkedlist: %s\n expected: 2, got: %d ", linkedlist, l)
	}
	if linkedlist.Head.Next.Data != 3 {
		t.Errorf("linkedlist data did not matched. linkedlist: %s\n expected: 3, got: %d ", linkedlist, linkedlist.Head.Next.Data)
	}
	// // testing insert in between
	linkedlist.InsertAt(1, 2)
	l = linkedlist.Length()
	if l != 3 {
		t.Fatalf("Length did not matched. linkedlist: %s\n expected: 3, got: %d ", linkedlist, l)
	}
	if linkedlist.Head.Next.Data != 2 {
		t.Errorf("linkedlist data did not matched. linkedlist: %s\n expected: 2, got: %d ", linkedlist, linkedlist.Head.Next.Data)
	}

	// testing insert in between
	linkedlist.InsertAt(2, 4)
	l = linkedlist.Length()
	if l != 4 {
		t.Fatalf("Length did not matched. linkedlist: %s\n expected: 3, got: %d ", linkedlist, l)
	}
	if linkedlist.Head.Next.Next.Data != 4 {
		t.Errorf("linkedlist data did not matched. linkedlist: %s\n expected: 4, got: %d ", linkedlist, linkedlist.Head.Next.Next.Data)
	}
}

func TestDeleteAtStart(t *testing.T) {
	var linkedlist SinglyLinkedList
	linkedlist.InsertAtEnd(1)
	linkedlist.InsertAtEnd(2)
	linkedlist.InsertAtEnd(3)
	l := linkedlist.Length()
	if l != 3 {
		t.Fatalf("Length did not matched. linkedlist: %s\n expected: 3, got: %d ", linkedlist, l)
	}
	if linkedlist.Head.Data != 1 {
		t.Errorf("linkedlist data did not matched. linkedlist: %s\n expected: 1, got: %d ", linkedlist, linkedlist.Head.Data)
	}

	linkedlist.DeleteAtStart()
	l = linkedlist.Length()
	if l != 2 {
		t.Fatalf("Length did not matched. linkedlist: %s\n expected: 2, got: %d ", linkedlist, l)
	}
	if linkedlist.Head.Data != 2 {
		t.Errorf("linkedlist data did not matched. linkedlist: %s\n expected: 2, got: %d ", linkedlist, linkedlist.Head.Data)
	}

	linkedlist.DeleteAtStart()
	l = linkedlist.Length()
	if l != 1 {
		t.Fatalf("Length did not matched. linkedlist: %s\n expected: 1, got: %d ", linkedlist, l)
	}
	if linkedlist.Head.Data != 3 {
		t.Errorf("linkedlist data did not matched. linkedlist: %s\n expected: 3, got: %d ", linkedlist, linkedlist.Head.Data)
	}

	linkedlist.DeleteAtStart()
	l = linkedlist.Length()
	if l != 0 {
		t.Fatalf("Length did not matched. linkedlist: %s\n expected: 0, got: %d ", linkedlist, l)
	}
	if linkedlist.Head != nil {
		t.Errorf("linkedlist data did not matched. linkedlist: %s\n expected: nil, got: %v ", linkedlist, linkedlist.Head)
	}
}

func TestDeleteAtEnd(t *testing.T) {
	var linkedlist SinglyLinkedList
	linkedlist.InsertAtEnd(1)
	linkedlist.InsertAtEnd(2)
	linkedlist.InsertAtEnd(3)
	l := linkedlist.Length()
	if l != 3 {
		t.Fatalf("Length did not matched. linkedlist: %s\n expected: 3, got: %d ", linkedlist, l)
	}
	if linkedlist.Head.Data != 1 {
		t.Errorf("linkedlist data did not matched. linkedlist: %s\n expected: 1, got: %d ", linkedlist, linkedlist.Head.Data)
	}

	linkedlist.DeleteAtEnd()
	l = linkedlist.Length()
	if l != 2 {
		t.Fatalf("Length did not matched. linkedlist: %s\n expected: 2, got: %d ", linkedlist, l)
	}
	if linkedlist.Head.Data != 1 && linkedlist.Head.Next.Data != 2 {
		t.Errorf("linkedlist data did not matched. linkedlist: %s\n expected: [1] --> [2], got: %d ", linkedlist, linkedlist.Head.Data)
	}

	linkedlist.DeleteAtEnd()
	l = linkedlist.Length()
	if l != 1 {
		t.Fatalf("Length did not matched. linkedlist: %s\n expected: 1, got: %d ", linkedlist, l)
	}
	if linkedlist.Head.Data != 1 {
		t.Errorf("linkedlist data did not matched. linkedlist: %s\n expected: 1, got: %d ", linkedlist, linkedlist.Head.Data)
	}

	linkedlist.DeleteAtEnd()
	l = linkedlist.Length()
	if l != 0 {
		t.Fatalf("Length did not matched. linkedlist: %s\n expected: 0, got: %d ", linkedlist, l)
	}
	if linkedlist.Head != nil {
		t.Errorf("linkedlist data did not matched. linkedlist: %s\n expected: nil, got: %v ", linkedlist, linkedlist.Head)
	}
}

func TestDeleteAtIndex(t *testing.T) {
	var linkedlist SinglyLinkedList
	linkedlist.InsertAtEnd(1)
	linkedlist.InsertAtEnd(2)
	linkedlist.InsertAtEnd(3)
	linkedlist.InsertAtEnd(4)
	l := linkedlist.Length()
	if l != 4 {
		t.Fatalf("Length did not matched. linkedlist: %s\n expected: 4, got: %d ", linkedlist, l)
	}
	if linkedlist.Head.Data != 1 {
		t.Errorf("linkedlist data did not matched. linkedlist: %s\n expected: 1, got: %d ", linkedlist, linkedlist.Head.Data)
	}

	linkedlist.DeleteAt(2)
	l = linkedlist.Length()
	if l != 3 {
		t.Fatalf("Length did not matched. linkedlist: %s\n expected: 3, got: %d ", linkedlist, l)
	}
	if linkedlist.Head.Data != 1 && linkedlist.Head.Next.Data != 2 && linkedlist.Head.Next.Next.Data != 4 {
		t.Errorf("linkedlist data did not matched. expected: 1 --> 2 --> 4 , got: %s ", linkedlist)
	}

	linkedlist.DeleteAt(2)
	l = linkedlist.Length()
	if l != 2 {
		t.Fatalf("Length did not matched. linkedlist: %s\n expected: 2, got: %d ", linkedlist, l)
	}
	if linkedlist.Head.Data != 1 && linkedlist.Head.Next.Data != 2 {
		t.Errorf("linkedlist data did not matched. expected: 1 --> 2 , got: %s ", linkedlist)
	}

	linkedlist.DeleteAt(0)
	l = linkedlist.Length()
	if l != 1 {
		t.Fatalf("Length did not matched. linkedlist: %s\n expected: 1, got: %d ", linkedlist, l)
	}
	if linkedlist.Head.Data != 2 {
		t.Errorf("linkedlist data did not matched. expected: 2, got: %d ", linkedlist.Head.Data)
	}
}

func TestIndexOf(t *testing.T) {
	var linkedlist SinglyLinkedList
	linkedlist.InsertAtEnd(10)
	linkedlist.InsertAtEnd(20)
	linkedlist.InsertAtEnd(30)
	linkedlist.InsertAtEnd(40)
	l := linkedlist.Length()
	if l != 4 {
		t.Fatalf("Length did not matched. linkedlist: %s\n expected: 4, got: %d ", linkedlist, l)
	}

	index := linkedlist.IndexOf(30)
	if index != 2 {
		t.Errorf("Index does not match. linkedlist: %s\n expected: 2, got: %d ", linkedlist, index)
	}

	index = linkedlist.IndexOf(33)
	if index != -1 {
		t.Errorf("Index does not match. linkedlist: %s\n expected: -1, got: %d ", linkedlist, index)
	}
}

func TestContains(t *testing.T) {
	var linkedlist SinglyLinkedList
	linkedlist.InsertAtEnd(10)
	linkedlist.InsertAtEnd(20)
	linkedlist.InsertAtEnd(30)
	linkedlist.InsertAtEnd(40)
	l := linkedlist.Length()
	if l != 4 {
		t.Fatalf("Length did not matched. linkedlist: %s\n expected: 4, got: %d ", linkedlist, l)
	}

	contains := linkedlist.Contains(30)
	if contains != true {
		t.Errorf("contains does not match. linkedlist: %s\n expected: true, got: %v ", linkedlist, contains)
	}

	contains = linkedlist.Contains(33)
	if contains != false {
		t.Errorf("contains does not match. linkedlist: %s\n expected: false, got: %v ", linkedlist, contains)
	}
}
