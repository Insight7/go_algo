package search

//BinarySearch implementations
func BinarySearch(arr []int, value int) int {
	return binarySearch(arr, value, 0, len(arr))
}

func binarySearch(arr []int, value int, left int, right int) int {
	if right >= left {
		mid := left + (right-left)/2
		if arr[mid] == value {
			return mid
		} else if arr[mid] > value {
			// left search
			return binarySearch(arr, value, left, mid-1)
		} else {
			// right search
			return binarySearch(arr, value, mid+1, right)
		}
	} else {
		return -1
	}
}
