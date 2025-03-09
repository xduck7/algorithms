package sliding_window

import (
	"fmt"
	"testing"
)

func TestFindLenOfUniqSubstring(t *testing.T) {
	tests := []struct {
		input    []rune
		expected int64
	}{
		{
			input:    []rune("abcabcbb"),
			expected: 3,
		},
		{
			input:    []rune("abcdef"),
			expected: 6,
		},
		{
			input:    []rune("aaaaaa"),
			expected: 1,
		},
		{
			input:    []rune(""),
			expected: 0,
		},
		{
			input:    []rune("a"),
			expected: 1,
		},
		{
			input:    []rune("pwwkew"),
			expected: 3,
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprint(idx), func(t *testing.T) {
			result := findLenOfUniqSubstring(tt.input)
			if result != tt.expected {
				t.Errorf("findLenOfUniqSubstring() = %v, want %v", result, tt.expected)
			}
		})
	}
}
