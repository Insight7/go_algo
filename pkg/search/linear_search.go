package search

//UnOrderedLinearSearch implementation
// search the given value in a unsorted array
func UnOrderedLinearSearch(arr []int, value int) int {
	for i := 0; i <= len(arr); i++ {
		if arr[i] == value {
			return i
		}
	}
	return -1
}

//OrderedLinearSearch implementation
//search the given value in a sorted array
func OrderedLinearSearch(arr []int, value int) int {
	for i := 0; i <= len(arr); i++ {
		if arr[i] == value {
			return i
		} else if arr[i] > value {
			return -1
		}
	}
	return -1
}
