package first

import (
	"testing"
)

func TestFirstLowerN(t *testing.T) {
	tests := []struct {
		arr      []int
		x        int
		expected int
	}{
		{[]int{1, 3, 5, 7, 9}, 6, 3},
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{10, 20, 30, 40}, 5, 0},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
		{[]int{}, 1, -1},
	}

	for _, test := range tests {
		result := firstNotLowerN(test.arr, test.x)
		if result != test.expected {
			t.Errorf("Alg(%v, %d) = %d; expected %d", test.arr, test.x, result, test.expected)
		}
	}
}
