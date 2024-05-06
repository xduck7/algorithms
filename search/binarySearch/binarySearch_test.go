package binarySearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		list     []int
		target   int
		expected bool
	}{
		{
			list:     []int{1, 2, 3, 4, 5, 23, 34, 45, 56, 67, 69, 143, 336, 475, 575, 676},
			target:   56,
			expected: true,
		},
		{
			list:     []int{1, 2, 3, 4, 5, 23, 34, 45, 56, 67, 69, 143, 336, 475, 575, 676},
			target:   6,
			expected: false,
		},
		{
			list:     []int{},
			target:   1,
			expected: false,
		},
		{
			list:     []int{1, 2, 3, 4, 5, 23, 34, 45, 56, 67, 69, 143, 336, 475, 575, 676},
			target:   1,
			expected: true,
		},
		{
			list:     []int{1, 2, 3, 4, 5, 23, 34, 45, 56, 67, 69, 143, 336, 475, 575, 676},
			target:   575,
			expected: true,
		},
	}

	for _, tc := range testCases {
		result := binarySearch(tc.list, tc.target)
		if result != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, result)
		}
	}
}
