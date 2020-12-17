package sorting

import (
	"math"
)

//BucketSort implementation in go
func BucketSort(data []int, bucketSize int) []int {

	// get minimin and maximum value for the array
	min, max := getMinMax(data)

	// create buckets
	nBuckets := int(max-min)/bucketSize + 1
	buckets := make([][]int, nBuckets)
	for i := 0; i < nBuckets; i++ {
		buckets[i] = make([]int, 0)
	}

	// put values in the appropriate buckets
	for _, n := range data {
		idx := int(n-min) / bucketSize
		buckets[idx] = append(buckets[idx], n)
	}
	// sort each bucket and combine them
	sorted := make([]int, 0)
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			BubbleSort(bucket)
			sorted = append(sorted, bucket...)
		}
	}
	// return the sorted array
	return sorted
}

func getMinMax(array []int) (minValue int, maxValue int) {
	maxValue = math.MinInt64
	minValue = 0
	for _, ele := range array {
		if ele > maxValue {
			maxValue = ele
		}
		if ele < minValue {
			minValue = ele
		}
	}
	return minValue, maxValue
}
