package nearestForX

import (
	"fmt"
	"testing"
)

func TestNearestForX(t *testing.T) {
	tests := []struct {
		arr      []int64
		x        int64
		expected int64
	}{
		{
			arr:      []int64{1, 3, 5, 7, 9},
			x:        5,
			expected: 2,
		},
		{
			arr:      []int64{1, 3, 5, 7, 9},
			x:        4,
			expected: 1,
		},
		{
			arr:      []int64{1, 3, 5, 7, 9},
			x:        8,
			expected: 3,
		},
		{
			arr:      []int64{1, 3, 5, 7, 9},
			x:        6,
			expected: 2,
		},
		{
			arr:      []int64{5, 5, 5, 5},
			x:        5,
			expected: 0,
		},
		{
			arr:      []int64{10},
			x:        10,
			expected: 0,
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprint(idx+1), func(t *testing.T) {
			result := nearestForX(tt.arr, tt.x)
			if result != tt.expected {
				t.Errorf("nearestForN() = %v, want %v", result, tt.expected)
			}
		})
	}
}
