package stack

//CapacityStack declaration
type CapacityStack struct {
	array    []int
	size     int
	capacity int
}

// MakeCapacityStack is a constructor for CapacityStack
func MakeCapacityStack(capacity int) CapacityStack {
	stack := CapacityStack{
		array:    make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
	return stack
}

//IsEmpty return true if stack is empty else false
func (stack *CapacityStack) IsEmpty() bool {
	return stack.size == 0
}

// Size returns size of the stack
func (stack *CapacityStack) Size() int {
	return stack.size
}

//Push adds a element at the end of the stack
func (stack *CapacityStack) Push(data int) {
	if stack.IsFull() {
		panic("Stack is full")
	}
	stack.array[stack.size] = data
	stack.size++
}

//Pop removes element from the end of the stack
func (stack *CapacityStack) Pop() {
	if stack.IsEmpty() {
		panic("Stack is empty")
	}
	stack.array = stack.array[:stack.size-1]
	stack.size--
}

//Peek returns the last element without removing it
func (stack *CapacityStack) Peek() int {
	if stack.IsEmpty() {
		panic("Stack is empty")
	}
	return stack.array[stack.size-1]
}

//IsFull returns true if the stack is full else returns false
func (stack *CapacityStack) IsFull() bool {
	return stack.size >= stack.capacity
}
