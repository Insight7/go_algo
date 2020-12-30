package queue

import (
	"testing"
)

func TestCapacityQueueEnqueue(t *testing.T) {
	queue := MakeCapacityQueue(5)
	queue.Enqueue(1)
	if queue.size != 1 {
		t.Errorf("Queue size does not match. expected 1, got %d", queue.size)
	}
	if queue.PeekFront() != 1 || queue.PeekBack() != 1 {
		t.Errorf("Queue element does not match. expected [1], got %v", queue.array)
	}

	queue.Enqueue(2)
	if queue.size != 2 {
		t.Errorf("Queue size does not match. expected 2, got %d", queue.size)
	}
	if queue.PeekFront() != 1 || queue.PeekBack() != 2 {
		t.Errorf("Queue element does not match. expected [1 2], got %v", queue.array)
	}
}

func TestCapacityQueueDequeue(t *testing.T) {
	queue := MakeCapacityQueue(5)
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
	if queue.PeekFront() != 2 || queue.PeekBack() != 3 {
		t.Errorf("Queue element does not match. expected [2 3], got %v", queue.array)
	}

	queue.Dequeue()
	if queue.size != 1 {
		t.Errorf("Queue size does not match. expected 2, got %d", queue.size)
	}
	if queue.PeekFront() != 3 || queue.PeekBack() != 3 {
		t.Errorf("Queue element does not match. expected [1], got %s", queue.String())
	}

	queue.Dequeue()
	if queue.size != 0 {
		t.Errorf("Queue size does not match. expected 2, got %d", queue.size)
	}
}

func TestCapacityQueuePeek(t *testing.T) {
	queue := MakeCapacityQueue(5)
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
