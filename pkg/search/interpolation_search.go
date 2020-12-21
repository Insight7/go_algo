package search

// InterpolationSearch caller
func InterpolationSearch(arr []int, value int) int {

	return interpolationSearch(arr, value, 0, len(arr))

}

// InterpolationSearch implementation
func interpolationSearch(arr []int, value int, left int, right int) int {
	if right >= left {
		slop := (right-left)/arr[right] - arr[left]
		pos := left + slop*(value-arr[left])

		if arr[pos] == value {
			return pos
		} else if arr[pos] > value {
			// left search
			return interpolationSearch(arr, value, left, pos-1)
		} else {
			// right search
			return interpolationSearch(arr, value, pos+1, right)
		}
	} else {
		return -1
	}
}
