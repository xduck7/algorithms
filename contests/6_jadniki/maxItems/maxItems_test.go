package maxitems_test

import (
	maxitems "algotithms/contests/6_jadniki/maxItems"
	"fmt"
	"testing"
)

func TestMaxItems(t *testing.T) {
	tests := []struct {
		nks      string
		prices   string
		expected int
	}{
		{"5 10", "4 2 1 10 5", 3},
		{"3 6", "2 2 2", 3},
		{"4 1", "2 2 2 2", 0},
	}

	for i, tt := range tests {
		result := maxitems.MaxItems(tt.nks, tt.prices)
		if result != tt.expected {
			t.Errorf("test case %d failed: expected %d, got %d", i+1, tt.expected, result)
		}
		fmt.Printf("#%d: OK\n", i+1)
	}
}
