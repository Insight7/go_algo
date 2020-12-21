package sorting

//MergeSort implementation in go
func MergeSort(array []int) []int {
	n := len(array)
	if n <= 1 {
		return array
	}
	left := MergeSort(array[0 : n/2])
	right := MergeSort(array[n/2:])

	return Merge(left, right)
}

// Merge two sorted arrays
func Merge(arr1 []int, arr2 []int) []int {
	n1 := len(arr1)
	n2 := len(arr2)
	totalLength := n1 + n2
	mergedArr := make([]int, totalLength)
	i1 := 0
	i2 := 0
	for i := 0; i < totalLength; i++ {
		if i1 == n1 {
			mergedArr[i] = arr2[i2]
			i2++
		} else if i2 == n2 {
			mergedArr[i] = arr1[i1]
			i1++
		} else {
			if arr1[i1] < arr2[i2] {
				mergedArr[i] = arr1[i1]
				i1++
			} else {
				mergedArr[i] = arr2[i2]
				i2++
			}
		}
	}
	return mergedArr

}
