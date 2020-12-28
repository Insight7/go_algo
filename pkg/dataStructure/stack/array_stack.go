package stack

//ArrayStack declaration
type ArrayStack struct {
	array []int
	size  int
}

//IsEmpty return true if stack is empty else false
func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}

// Size returns size of the stack
func (stack *ArrayStack) Size() int {
	return stack.size
}

//Push adds a element at the end of the stack
func (stack *ArrayStack) Push(data int) {
	stack.array = append(stack.array, data)
	stack.size++
}

//Pop removes element from the top of the stack
func (stack *ArrayStack) Pop() {
	if stack.IsEmpty() {
		panic("Stack is empty")
	}
	stack.array = stack.array[:stack.size-1]
	stack.size--
}

//Peek returns the last element without removing it
func (stack *ArrayStack) Peek() int {
	if stack.IsEmpty() {
		panic("Stack is empty")
	}
	return stack.array[stack.size-1]
}
