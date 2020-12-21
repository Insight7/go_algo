package sorting

// SelectionSort implementation in go
func SelectionSort(array []int) {
	var minIndex int
	for i := 0; i < len(array); i++ {
		minIndex = findMinValueIndex(i, array)
		Swap(&array[i], &array[minIndex])
	}
}

func findMinValueIndex(startIndex int, arr []int) int {
	minValueIndex := startIndex
	for i := startIndex; i < len(arr); i++ {
		if arr[minValueIndex] > arr[i] {
			minValueIndex = i
		}
	}
	return minValueIndex

}
