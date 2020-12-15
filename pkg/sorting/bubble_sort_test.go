package sorting

import (
	"testing"
)

var flagTests = []struct {
	expected []int
	input    []int
}{
	{[]int{}, []int{}},
	{[]int{1}, []int{1}},
	{[]int{1, 2}, []int{1, 2}},
	{[]int{1, 2}, []int{2, 1}},
	{[]int{1, 1}, []int{1, 1}},

	{[]int{1, 2, 3}, []int{1, 2, 3}},
	{[]int{1, 2, 3}, []int{1, 3, 2}},
	{[]int{1, 2, 3}, []int{2, 1, 3}},
	{[]int{1, 2, 3}, []int{2, 3, 1}},
	{[]int{1, 2, 3}, []int{3, 1, 2}},
	{[]int{1, 2, 3}, []int{3, 2, 1}},

	{[]int{1, 1, 3}, []int{1, 1, 3}},
	{[]int{1, 1, 3}, []int{1, 3, 1}},
	{[]int{1, 1, 3}, []int{3, 1, 1}},

	{[]int{1, 1, 1}, []int{1, 1, 1}},
}

func TestBubbleSort(t *testing.T) {
	t.Parallel()
	for _, value := range flagTests {
		t.Run("", func(t *testing.T) {
			t.Parallel()
			result := BubbleSort(value.input)
			if !IsArrayEqual(result, value.expected) {
				t.Errorf("soriting failed expected:  %#v \n output: %#v \n", value.expected, value.input)
			}

		})

	}
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
