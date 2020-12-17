package sorting

//InsertionSort implementation in go
func InsertionSort(array []int) {
	n := len(array)
	if n == 0 || n == 1 {
		return
	}
	var j int
	for i := 1; i < n; i++ {
		// i is the marker before which
		unsortedElement := array[i]
		j = i - 1
		for j >= 0 && array[j] > unsortedElement {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = unsortedElement
	}

}
