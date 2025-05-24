package maxway_test

import (
	maxway "algotithms/contests/7_dynamic/maxWay"
	"fmt"
	"testing"
)

func TestMaxSumPath(t *testing.T) {
	tests := []struct {
		n        int
		str      string
		expected int
	}{
		{5, "1 2 3 4 5", 15},
		{3, "10 -5 20", 30},
		{1, "7", 7},
		{4, "-1 -2 -3 -4", -6},
		{6, "5 10 20 15 25 10", 85},
	}

	for i, tt := range tests {
		res := maxway.MaxWay(tt.n, tt.str)
		if res != tt.expected {
			t.Errorf("test case %d failed: expected %d, got %d", i+1, tt.expected, res)
		}
		fmt.Printf("#%d: OK\n", i+1)
	}
}
