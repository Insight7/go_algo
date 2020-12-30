package queue

import "fmt"

//CapacityQueue declaration
type CapacityQueue struct {
	array    []int
	size     int
	capacity int
	front    int
	back     int
}

// MakeCapacityQueue is a constructor for CapacityQueue
func MakeCapacityQueue(capacity int) CapacityQueue {
	queue := CapacityQueue{
		array:    make([]int, capacity),
		size:     0,
		front:    0,
		back:     0,
		capacity: capacity,
	}
	return queue
}

//IsEmpty return true if queue is empty else false
func (queue *CapacityQueue) IsEmpty() bool {
	return queue.size == 0
}

// Size returns size of the queue
func (queue *CapacityQueue) Size() int {
	return queue.size
}

//Enqueue add a element at the end of a queue
func (queue *CapacityQueue) Enqueue(data int) {
	if queue.IsFull() {
		panic("Queue is full")
	}
	queue.back++
	queue.back = queue.back % queue.capacity
	queue.array[queue.back] = data
	queue.size++
}

//Dequeue remove and returns the element from the front of the queue
func (queue *CapacityQueue) Dequeue() int {
	if queue.IsEmpty() {
		panic("Queue is Empty")
	}
	data := queue.array[queue.front]
	queue.front++
	queue.front = queue.front % queue.capacity
	queue.size--
	return data
}

//PeekFront returns the front element without removing it
func (queue *CapacityQueue) PeekFront() int {
	if queue.IsEmpty() {
		panic("Queue is empty")
	}
	return queue.array[queue.front+1]
}

//PeekBack returns the back element without removing it
func (queue *CapacityQueue) PeekBack() int {
	if queue.IsEmpty() {
		panic("Queue is empty")
	}
	return queue.array[queue.back]
}

//IsFull returns true if the stack is full else returns false
func (queue *CapacityQueue) IsFull() bool {
	return queue.size >= queue.capacity
}

func (queue *CapacityQueue) String() string {
	return fmt.Sprintf("CapacityQueue{ array: %v, size: %d, front :%d, back: %d }", queue.array, queue.size, queue.front, queue.back)
}
