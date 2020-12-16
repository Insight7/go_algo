package sorting

//BubbleSort implementation in go
func BubbleSort(array []int) {
	if array == nil {
		return
	}

	n := len(array)
	var isSwappingDone bool
	for i := 0; i < n-1; i++ {
		isSwappingDone = false
		for j := 0; j < n-i-1; j++ {
			// swap the bigger value to the end so that at the end of every iteration the maximum value is at the end
			if array[j] > array[j+1] {
				Swap(&array[j], &array[j+1])
				isSwappingDone = true
			}
		}
		if !isSwappingDone {
			break
		}
	}

	// The slice in go are just a reference
	// hence changing the value here changes the value of the input array tpp
}

//Swap used to swap two values
func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
