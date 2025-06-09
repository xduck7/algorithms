package taskRewards_test

import (
	"algotithms/contests/6_jadniki/taskRewards"
	"fmt"
	"testing"
)

func TestMaxTotalReward(t *testing.T) {
	tests := []struct {
		n        int
		dw       string
		expected int
	}{
		{5, "1 10\n2 20\n2 30\n1 40\n2 50", 90},
		{3, "1 100\n1 200\n1 50", 200},
		{4, "4 70\n2 60\n4 50\n3 40", 220},
	}
	for i, tt := range tests {
		result := taskRewards.MaxTotalReward(tt.n, tt.dw)
		if result != tt.expected {
			t.Errorf("test case %d failed: expected %d, got %d", i+1, tt.expected, result)
		}
		fmt.Printf("#%d: OK\n", i+1)
	}
}
