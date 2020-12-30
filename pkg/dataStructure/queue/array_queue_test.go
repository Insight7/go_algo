package queue

import (
	"testing"
)

func TestArrayQueueEnqueue(t *testing.T) {
	queue := ArrayQueue{}
	queue.Enqueue(1)
	if queue.size != 1 {
		t.Errorf("Queue size does not match. expected 1, got %d", queue.size)
	}
	if queue.array[0] != 1 {
		t.Errorf("Queue element does not match. expected [1], got %v", queue.array)
	}

	queue.Enqueue(2)
	if queue.size != 2 {
		t.Errorf("Queue size does not match. expected 2, got %d", queue.size)
	}
	if queue.array[0] != 1 || queue.array[1] != 2 {
		t.Errorf("Queue element does not match. expected [1 2], got %v", queue.array)
	}
}

func TestArrayQueueDequeue(t *testing.T) {
	queue := ArrayQueue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	if queue.size != 3 {
		t.Errorf("Queue size does not match. expected 3, got %d", queue.size)
	}

	queue.Dequeue()
	if queue.size != 2 {
		t.Errorf("Queue size does not match. expected 2, got %d", queue.size)
	}
	if queue.array[0] != 2 || queue.array[1] != 3 {
		t.Errorf("Queue element does not match. expected [2 3], got %v", queue.array)
	}

	queue.Dequeue()
	if queue.size != 1 {
		t.Errorf("Queue size does not match. expected 2, got %d", queue.size)
	}
	if queue.array[0] != 3 || len(queue.array) != 1 {
		t.Errorf("Queue element does not match. expected [1], got %v", queue.array)
	}

	queue.Dequeue()
	if queue.size != 0 {
		t.Errorf("Queue size does not match. expected 2, got %d", queue.size)
	}
}

func TestArrayQueuePeek(t *testing.T) {
	queue := ArrayQueue{}
	queue.Enqueue(1)
	if queue.size != 1 {
		t.Errorf("Queue size does not match. expected 1, got %d", queue.size)
	}
	if queue.PeekFront() != 1 || queue.PeekBack() != 1 {
		t.Errorf("Queue element does not match. expected 1, got %v", queue.array)
	}

	queue.Enqueue(2)
	if queue.size != 2 {
		t.Errorf("Queue size does not match. expected 2, got %d", queue.size)
	}
	if queue.PeekFront() != 1 || queue.PeekBack() != 2 {
		t.Errorf("Queue element does not match. expected [1 2], got %v", queue.array)
	}

	queue.Enqueue(3)
	if queue.size != 3 {
		t.Errorf("Queue size does not match. expected 3, got %d", queue.size)
	}
	if queue.PeekFront() != 1 || queue.PeekBack() != 3 {
		t.Errorf("Queue element does not match. expected [1 3], got %v", queue.array)
	}
}
