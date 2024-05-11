package selectionSort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	testCases := []struct {
		list     []int
		expected []int
	}{
		{
			list:     []int{34, 1, 575, 475, 336, 2, 5, 23, 56, 143, 676, 45, 3, 4, 67, 69},
			expected: []int{1, 2, 3, 4, 5, 23, 34, 45, 56, 67, 69, 143, 336, 475, 575, 676},
		},
	}

	for _, tc := range testCases {
		result := selectionSort(tc.list)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}
