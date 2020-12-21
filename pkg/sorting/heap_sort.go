package sorting

//HeapSort implementation in go
func HeapSort(array []int) {
	n := len(array)
	// build a heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(array, i, n)
	}
	// extract min and hepify the array again
	for i := n - 1; i > 0; i-- {
		Swap(&array[0], &array[i])
		heapify(array, 0, i)
	}
}

func heapify(arr []int, i int, n int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		Swap(&arr[i], &arr[largest])
		heapify(arr, largest, n)
	}
}
