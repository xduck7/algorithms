package ctwobackpacks_test

import (
	"testing"

	ctwobackpacks "algotithms/contests/8_two-dem_dynamic/c_two_backpacks"
)

func TestMaxTotalValue(t *testing.T) {
	type testCase struct {
		n, w1, w2 int64
		items     [][2]int64
		expected  int64
	}

	tests := []testCase{
		{
			n:  3,
			w1: 4,
			w2: 5,
			items: [][2]int64{
				{2, 3},
				{3, 4},
				{4, 5},
			},
			expected: 12,
		},
		{
			n:  2,
			w1: 5,
			w2: 5,
			items: [][2]int64{
				{5, 10},
				{5, 20},
			},
			expected: 30,
		},
		{
			n:  4,
			w1: 5,
			w2: 5,
			items: [][2]int64{
				{2, 3},
				{2, 4},
				{3, 5},
				{4, 6},
			},
			expected: 15,
		},
	}

	for idx, tc := range tests {
		result := ctwobackpacks.TwoBackpacks(tc.n, tc.w1, tc.w2, tc.items)
		if result != tc.expected {
			t.Errorf("Test %d failed: expected %d, got %d", idx+1, tc.expected, result)
		}
	}
}
