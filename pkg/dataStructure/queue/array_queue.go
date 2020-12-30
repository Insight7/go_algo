package queue

//ArrayQueue declaration
type ArrayQueue struct {
	array []int
	size  int
}

//IsEmpty return true if queue is empty else false
func (queue *ArrayQueue) IsEmpty() bool {
	return queue.size == 0
}

// Size returns size of the queue
func (queue *ArrayQueue) Size() int {
	return queue.size
}

//Enqueue add a element at the end of a queue
func (queue *ArrayQueue) Enqueue(data int) {
	queue.array = append(queue.array, data)
	queue.size++
}

//Dequeue remove and returns the element from the front of the queue
func (queue *ArrayQueue) Dequeue() int {
	if queue.IsEmpty() {
		panic("Queue is empty")
	}
	data := queue.array[0]
	queue.array = queue.array[1:]
	queue.size--
	return data

}

//PeekFront returns the front element without removing it
func (queue *ArrayQueue) PeekFront() int {
	if queue.IsEmpty() {
		panic("Queue is empty")
	}
	return queue.array[0]

}

//PeekBack returns the back element without removing it
func (queue *ArrayQueue) PeekBack() int {
	if queue.IsEmpty() {
		panic("Queue is empty")
	}
	return queue.array[queue.size-1]
}
