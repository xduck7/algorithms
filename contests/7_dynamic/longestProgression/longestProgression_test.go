package longestprogression_test

import (
	longestprogression "algotithms/contests/7_dynamic/longestProgression"
	"fmt"
	"testing"
)

func TestLongestProgression(t *testing.T) {
	tests := []struct {
		n        int
		nums     string
		expected int
	}{
		{5, "1 2 3 4 5", 5},
		{5, "5 4 3 2 1", 1},
		{8, "10 22 9 33 21 50 41 60", 5},
		{1, "7", 1},
	}

	for i, tt := range tests {
		result := longestprogression.LongestProgression(tt.n, tt.nums)
		if result != tt.expected {
			t.Errorf("test case %d failed: expected %d, got %d", i+1, tt.expected, result)
		}
		fmt.Printf("#%d: OK\n", i+1)
	}
}
