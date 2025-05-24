package countsplitting_test

import (
	countsplitting "algotithms/contests/7_dynamic/countSplitting"
	"fmt"
	"testing"
)

func TestCountPartitions(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 7},
		{10, 42},
	}
	for i, tt := range tests {
		got := countsplitting.CountSplitting(tt.n)
		if got != tt.expected {
			t.Errorf("res: %d)= %d, want %d", tt.n, got, tt.expected)
		}
		fmt.Printf("#%d: OK\n", i+1)
	}
}
