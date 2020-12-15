package sorting

//BubbleSort implementation in go
func BubbleSort(array []int) []int {
	if array == nil {
		return nil
	}

	n := len(array)
	var isSwappingDone bool

	for i := 0; i < n-1; i++ {

		isSwappingDone = false
		for j := 0; j < n-i-1; j++ {
			if array[j] > array[j+1] {
				Swap(&array[j], &array[j+1])
				isSwappingDone = true
			}
		}
		if !isSwappingDone {
			break
		}
	}

	return array

}

//Swap used to swap two values
func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
