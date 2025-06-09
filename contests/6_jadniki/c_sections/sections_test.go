package sections_test

import (
	"algotithms/contests/6_jadniki/sections"
	"fmt"
	"testing"
)

func TestMaxSectionStrick(t *testing.T) {
	tests := []struct {
		n        int
		sec      string
		expected int
	}{
		{5, "1 3\n2 5\n4 6\n6 7\n5 9", 3},
		{3, "1 2\n2 3\n3 4", 3},
		{4, "1 10\n2 3\n3 4\n5 6", 3},
	}
	for i, tt := range tests {
		result := sections.MaxSectionStrick(tt.n, tt.sec)
		if result != tt.expected {
			t.Errorf("test case %d failed: expected %d, got %d", i+1, tt.expected, result)
		}
		fmt.Printf("#%d: OK\n", i+1)
	}
}
