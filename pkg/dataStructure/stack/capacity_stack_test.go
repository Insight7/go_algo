package stack

import (
	"testing"
)

func TestCapacityStackPush(t *testing.T) {
	stack := MakeCapacityStack(5)
	stack.Push(1)
	if stack.size != 1 {
		t.Errorf("Stack size does not match. expected 1, got %d", stack.size)
	}
	if stack.array[0] != 1 {
		t.Errorf("Stack element does not match. expected [1], got %v", stack.array)
	}

	stack.Push(2)
	if stack.size != 2 {
		t.Errorf("Stack size does not match. expected 2, got %d", stack.size)
	}
	if stack.array[0] != 1 || stack.array[1] != 2 {
		t.Errorf("Stack element does not match. expected [1 2], got %v", stack.array)
	}
}

func TestCapacityStackPop(t *testing.T) {
	stack := MakeCapacityStack(5)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.size != 3 {
		t.Errorf("Stack size does not match. expected 3, got %d", stack.size)
	}

	stack.Pop()
	if stack.size != 2 {
		t.Errorf("Stack size does not match. expected 2, got %d", stack.size)
	}
	if stack.array[0] != 1 || stack.array[1] != 2 {
		t.Errorf("Stack element does not match. expected [1 2], got %v", stack.array)
	}

	stack.Pop()
	if stack.size != 1 {
		t.Errorf("Stack size does not match. expected 2, got %d", stack.size)
	}
	if stack.array[0] != 1 || len(stack.array) != 1 {
		t.Errorf("Stack element does not match. expected [1], got %v", stack.array)
	}

	stack.Pop()
	if stack.size != 0 {
		t.Errorf("Stack size does not match. expected 2, got %d", stack.size)
	}
}

func TestCapacityStackPeek(t *testing.T) {
	stack := MakeCapacityStack(5)
	stack.Push(1)
	if stack.size != 1 {
		t.Errorf("Stack size does not match. expected 1, got %d", stack.size)
	}
	if stack.Peek() != 1 {
		t.Errorf("Stack element does not match. expected 1, got %v", stack.array)
	}

	stack.Push(2)
	if stack.size != 2 {
		t.Errorf("Stack size does not match. expected 2, got %d", stack.size)
	}
	if stack.Peek() != 2 {
		t.Errorf("Stack element does not match. expected [1 2], got %v", stack.array)
	}

	stack.Pop()
	if stack.size != 1 {
		t.Errorf("Stack size does not match. expected 2, got %d", stack.size)
	}
	if stack.Peek() != 1 {
		t.Errorf("Stack element does not match. expected 1, got %v", stack.array)
	}
}
