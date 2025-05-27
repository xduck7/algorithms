package maxInWindow_test

import (
	"algotithms/contests/9_data_struct/maxInWindow"
	"reflect"
	"testing"
)

func TestMaxInWindow(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1, 2, 3, 4, 5}, 2, []int{2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, 3, []int{5, 4, 3}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, -1}, 1, []int{1, -1}},
		{[]int{2, 1}, 2, []int{2}},
		{[]int{}, 3, []int{}},
		{[]int{4, 3, 5, 4, 3, 3, 6, 7}, 3, []int{5, 5, 5, 4, 6, 7}},
	}

	for _, tt := range tests {
		res := maxInWindow.MaxInWindow(tt.nums, tt.k)
		if !reflect.DeepEqual(res, tt.expected) {
			t.Errorf("MaxInSlidingWindow(%v, %d) = %v; expected %v", tt.nums, tt.k, res, tt.expected)
		}
	}
}
