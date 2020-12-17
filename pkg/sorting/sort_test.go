package sorting

import (
	"math/rand"
	"sort"
	"testing"
)

const RandomArraySize int = 1000
const RandomSortIterationCount int = 1000

// Bubble sort test cases
func TestBubbleSort(t *testing.T) {
	for i := 0; i < RandomSortIterationCount; i++ {
		randomArray := randomArrayGenerator(i)
		result := copySlice(randomArray)
		//testArray := copySlice(randomArray)

		BubbleSort(result)
		sort.Ints(randomArray)

		if !IsArrayEqual(result, randomArray) {
			//t.Logf("soriting failed.\ntestArray:\t\t\t%#v\nexpected:\t\t\t%#v\noutput:\t\t\t%#v \n", testArray, randomArray, result)
			t.Fail()
		}
	}

}

// Bucket sort test cases
func TestBucketSort(t *testing.T) {
	for i := 0; i < RandomSortIterationCount; i++ {
		randomArray := randomArrayGeneratorWithRange(i, 1000)
		result := copySlice(randomArray)
		//testArray := copySlice(randomArray)

		result = BucketSort(result, 10)
		sort.Ints(randomArray)

		if !IsArrayEqual(result, randomArray) {
			//t.Logf("soriting failed.\ntestArray:\t\t\t%#v\nexpected:\t\t\t%#v\noutput:\t\t\t%#v \n", testArray, randomArray, result)
			t.Fail()
		}

	}
}

// Counting sort test cases
func TestCountingSort(t *testing.T) {
	for i := 0; i < RandomSortIterationCount; i++ {
		randomArray := randomArrayGenerator(i)
		result := copySlice(randomArray)
		//testArray := copySlice(randomArray)

		CountingSort(result)
		sort.Ints(randomArray)

		if !IsArrayEqual(result, randomArray) {
			//t.Logf("soriting failed.\ntestArray:\t\t\t%#v\nexpected:\t\t\t%#v\noutput:\t\t\t%#v \n", testArray, randomArray, result)
			t.Fail()
		}

	}
}

// Heap sort test cases
func TestHeapSort(t *testing.T) {
	for i := 0; i < RandomSortIterationCount; i++ {
		randomArray := randomArrayGenerator(i)
		result := copySlice(randomArray)
		//testArray := copySlice(randomArray)

		HeapSort(result)
		sort.Ints(randomArray)

		if !IsArrayEqual(result, randomArray) {
			//t.Logf("soriting failed.\ntestArray:\t\t\t%#v\nexpected:\t\t\t%#v\noutput:\t\t\t%#v \n", testArray, randomArray, result)
			t.Fail()
		}

	}
}

// insertion sort test cases

func TestInsertionSort(t *testing.T) {
	for i := 0; i < RandomSortIterationCount; i++ {
		randomArray := randomArrayGenerator(i)
		result := copySlice(randomArray)
		//testArray := copySlice(randomArray)

		InsertionSort(result)
		sort.Ints(randomArray)

		if !IsArrayEqual(result, randomArray) {
			//t.Logf("soriting failed.\ntestArray:\t\t\t%#v\nexpected:\t\t\t%#v\noutput:\t\t\t%#v \n", testArray, randomArray, result)
			t.Fail()
		}

	}
}

// Merge sort test cases
func TestMergeSort(t *testing.T) {
	for i := 0; i < RandomSortIterationCount; i++ {
		randomArray := randomArrayGenerator(i)
		result := copySlice(randomArray)
		//testArray := copySlice(randomArray)

		MergeSort(result)
		sort.Ints(randomArray)

		if !IsArrayEqual(result, randomArray) {
			//t.Logf("soriting failed.\ntestArray:\t\t\t%#v\nexpected:\t\t\t%#v\noutput:\t\t\t%#v \n", testArray, randomArray, result)
			t.Fail()
		}

	}
}

// Selection sort test cases
func TestSelectionSort(t *testing.T) {
	for i := 0; i < RandomSortIterationCount; i++ {
		randomArray := randomArrayGenerator(i)
		result := copySlice(randomArray)
		//testArray := copySlice(randomArray)

		SelectionSort(result)
		sort.Ints(randomArray)

		if !IsArrayEqual(result, randomArray) {
			//t.Logf("soriting failed.\ntestArray:\t\t\t%#v\nexpected:\t\t\t%#v\noutput:\t\t\t%#v \n", testArray, randomArray, result)
			t.Fail()
		}

	}
}

// Radix sort test cases

func TestRadixSort(t *testing.T) {
	for i := 0; i < RandomSortIterationCount; i++ {
		randomArray := randomArrayGenerator(i)
		result := copySlice(randomArray)
		//testArray := copySlice(randomArray)

		RadixSort(result)
		sort.Ints(randomArray)

		if !IsArrayEqual(result, randomArray) {
			//t.Logf("soriting failed.\ntestArray:\t\t\t%#v\nexpected:\t\t\t%#v\noutput:\t\t\t%#v \n", testArray, randomArray, result)
			t.Fail()
		}

	}
}

// Test helper functions

func randomArrayGeneratorWithRange(seed int, valueRange int) []int {
	array := make([]int, RandomArraySize)
	rand.Seed(int64(seed))
	for i := 0; i < RandomArraySize; i++ {
		array[i] = int(rand.Uint32())
		array[i] = rand.Int() % valueRange
	}
	return array
}

func randomArrayGenerator(seed int) []int {
	return randomArrayGeneratorWithRange(seed, 100000)
}

func copySlice(input []int) []int {
	result := make([]int, len(input))
	copy(result, input)
	return result
}

func IsArrayEqual(input []int, output []int) bool {
	if len(input) != len(output) {
		return false
	}

	l := len(input)
	for i := 0; i < l; i++ {
		if input[i] != output[i] {
			return false
		}
	}
	return true
}
