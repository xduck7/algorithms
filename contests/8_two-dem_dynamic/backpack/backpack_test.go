package backpack_test

import (
	"algotithms/contests/8_two-dem_dynamic/backpack"
	"testing"
)

func TestBackpack(t *testing.T) {
	tests := []struct {
		n, w     int
		vm, wm   []int
		expected int
	}{
		{
			n:        4,
			w:        7,
			vm:       []int{6, 1, 2, 5},
			wm:       []int{13, 1, 6, 10},
			expected: 16,
		},
		{
			n:        3,
			w:        50,
			vm:       []int{10, 20, 30},
			wm:       []int{60, 100, 120},
			expected: 220,
		},
		{
			n:        1,
			w:        5,
			vm:       []int{6},
			wm:       []int{100},
			expected: 0,
		},
		{
			n:        2,
			w:        3,
			vm:       []int{1, 2},
			wm:       []int{2, 3},
			expected: 5,
		},
	}

	for _, tt := range tests {
		result := backpack.Backpack(tt.n, tt.w, tt.vm, tt.wm)
		if result != tt.expected {
			t.Errorf("backpack(%d, %d, %v, %v) = %d; expected %d",
				tt.n, tt.w, tt.vm, tt.wm, result, tt.expected)
		}
	}
}
