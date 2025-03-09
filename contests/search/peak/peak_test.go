package peak

import (
	"fmt"
	"testing"
)

func TestFindPeak(t *testing.T) {
	tests := []struct {
		arr      []int64
		expected int64
	}{
		{
			arr:      []int64{1, 3, 5, 4, 2},
			expected: 2,
		},
		{
			arr:      []int64{10, 8, 6, 5, 3},
			expected: 0,
		},
		{
			arr:      []int64{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			arr:      []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: 9,
		},
		{
			arr:      []int64{42},
			expected: 0,
		},
		{
			arr:      []int64{10, 5},
			expected: 0,
		},
		{
			arr:      []int64{5, 10},
			expected: 1,
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprint(idx), func(t *testing.T) {
			result := findPeak(tt.arr)
			if result != tt.expected {
				t.Errorf("findPeak() = %v, want %v", result, tt.expected)
			}
		})
	}
}
