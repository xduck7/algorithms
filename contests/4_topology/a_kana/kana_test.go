package kana_test

import (
	"algotithms/contests/4_topology/kana"
	"fmt"
	"testing"
)

func TestKana(t *testing.T) {
	tests := []struct {
		N        int
		v        [][]int
		expected []int
		hasCycle bool
	}{
		{
			N: 4,
			v: [][]int{
				{},
				{2, 3},
				{4},
				{4},
				{},
			},
			expected: []int{1, 2, 3, 4},
			hasCycle: false,
		},
		{
			N: 3,
			v: [][]int{
				{},
				{2},
				{3},
				{1},
			},
			expected: nil,
			hasCycle: true,
		},
		{
			N: 5,
			v: [][]int{
				{},
				{2},
				{},
				{4, 5},
				{},
				{},
			},
			expected: []int{1, 3, 2, 4, 5},
			hasCycle: false,
		},
		{
			N: 1,
			v: [][]int{
				{},
				{},
			},
			expected: []int{1},
			hasCycle: false,
		},
	}

	for id, tt := range tests {
		t.Run(fmt.Sprintf("%d", id), func(t *testing.T) {
			result := kana.Kana(tt.N, tt.v)

			if tt.hasCycle {
				if result != nil {
					t.Errorf("got topological order: %v", result)
				}
				return
			}

			if result == nil {
				t.Error("got cycle detection")
				return
			}

			if len(result) != tt.N {
				t.Errorf("wrong length: expected %d, got %d", tt.N, len(result))
			}

			pos := make(map[int]int)
			for i, v := range result {
				pos[v] = i
			}

			for u := 1; u <= tt.N; u++ {
				for _, v := range tt.v[u] {
					if pos[u] > pos[v] {
						t.Errorf("edge %d > %d but %d appears before %d", u, v, u, v)
					}
				}
			}
		})
	}
}
