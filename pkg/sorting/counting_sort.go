package sorting

//CountingSort implementation in go
// The array values range from 0 to maxValue
func CountingSort(array []int, maxValue int) {
	count := make([]int, maxValue)

	// count elements presence
	for _, ele := range array {
		count[ele]++
	}
	i := 0
	for index, value := range count {
		if value != 0 {
			for j := 0; j < value; j++ {
				array[i] = index
				i++
			}
		}
	}

}
