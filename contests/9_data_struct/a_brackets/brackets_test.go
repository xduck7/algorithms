package brackets_test

import (
	"algotithms/contests/9_data_struct/brackets"
	"testing"
)

func TestBracketsYesOrNo(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		{"", true},
		{"()", true},
		{"([]{})", true},
		{"({[]})", true},
		{"([)]", false},
		{"(", false},
		{")", false},
		{"{[()]}", true},
		{"{[(])}", false},
		{"((()))", true},
		{"((())", false},
	}

	for _, tt := range tests {
		result := brackets.BracketsYesOrNo(tt.s)
		if result != tt.expected {
			t.Errorf("brackets(%q) = %v; expected %v", tt.s, result, tt.expected)
		}
	}
}
