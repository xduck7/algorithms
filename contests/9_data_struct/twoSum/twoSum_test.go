package twoSum_test

import (
	"algotithms/contests/9_data_struct/twoSum"
	"testing"
)

func TestHasTwoSum(t *testing.T) {
	tests := []struct {
		a        []int
		k        int
		expected bool
	}{
		{[]int{1, 2, 3, 4}, 5, true},
		{[]int{1, 2, 3, 4}, 8, false},
		{[]int{-1, -2, -3, -4}, -5, true},
		{[]int{5, 75, 25}, 100, true},
		{[]int{1}, 2, false},
		{[]int{}, 0, false},
		{[]int{2, 2, 3}, 4, true},
	}

	for _, tt := range tests {
		res := twoSum.TwoSum(tt.a, tt.k)
		if res != tt.expected {
			t.Errorf("twoSum(%v, %d) = %v; expected %v", tt.a, tt.k, res, tt.expected)
		}
	}
}
