package nop_test

import (
	"algotithms/contests/8_two-dem_dynamic/nop"
	"testing"
)

func TestLCS(t *testing.T) {
	tests := []struct {
		a, b     string
		expected string
	}{
		{"ABACA", "BCBA", "BCA"},
		{"ABCBDAB", "BDCAB", "BDAB"},
		{"", "ABC", ""},
		{"ABC", "", ""},
		{"ABC", "DEF", ""},
		{"AGGTAB", "GXTXAYB", "GTAB"},
		{"ABCDEF", "FBDAMN", "BD"},
	}

	for _, tt := range tests {
		result := nop.Nop(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Nop(%q, %q) = %q; expected %q", tt.a, tt.b, result, tt.expected)
		}
	}
}
