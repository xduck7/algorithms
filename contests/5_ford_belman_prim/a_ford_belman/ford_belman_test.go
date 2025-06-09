package ford_belman

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFordBellman(t *testing.T) {
	tests := []struct {
		graph            func() *Graph
		start            int64
		end              int64
		expectedDist     int64
		expectedPath     []int64
		hasNegativeCycle bool
	}{
		{
			graph: func() *Graph {
				g := &Graph{}
				g.AddVertex(1)
				g.AddVertex(2)
				g.AddVertex(3)
				g.AddWeightedEdge(1, 2, 5)
				g.AddWeightedEdge(1, 3, 3)
				g.AddWeightedEdge(2, 3, 1)
				return g
			},
			start:            1,
			end:              3,
			expectedDist:     3,
			expectedPath:     []int64{1, 3},
			hasNegativeCycle: false,
		},
		{
			graph: func() *Graph {
				g := &Graph{}
				g.AddVertex(1)
				g.AddVertex(2)
				g.AddVertex(3)
				g.AddWeightedEdge(1, 2, 5)
				return g
			},
			start:            1,
			end:              3,
			expectedDist:     -1,
			expectedPath:     nil,
			hasNegativeCycle: false,
		},
		{
			graph: func() *Graph {
				g := &Graph{}
				g.AddVertex(1)
				g.AddVertex(2)
				g.AddVertex(3)
				g.AddVertex(4)
				g.AddWeightedEdge(1, 2, 1)
				g.AddWeightedEdge(2, 3, -2)
				g.AddWeightedEdge(3, 4, 2)
				g.AddWeightedEdge(1, 4, 5)
				return g
			},
			start:            1,
			end:              4,
			expectedDist:     1,
			expectedPath:     []int64{1, 2, 3, 4},
			hasNegativeCycle: false,
		},
		{
			graph: func() *Graph {
				g := &Graph{}
				g.AddVertex(1)
				g.AddVertex(2)
				g.AddVertex(3)
				g.AddWeightedEdge(1, 2, 1)
				g.AddWeightedEdge(2, 3, -2)
				g.AddWeightedEdge(3, 2, 1) // отрицательный цикл 2-3-2
				return g
			},
			start:            1,
			end:              3,
			expectedDist:     -1,
			expectedPath:     nil,
			hasNegativeCycle: true,
		},
		{
			graph: func() *Graph {
				g := &Graph{}
				g.AddVertex(1)
				g.AddVertex(2)
				g.AddWeightedEdge(1, 2, 5)
				return g
			},
			start:            1,
			end:              1,
			expectedDist:     0,
			expectedPath:     []int64{1},
			hasNegativeCycle: false,
		},
	}

	for id, tt := range tests {
		t.Run(fmt.Sprintf("%d", id+1), func(t *testing.T) {
			g := tt.graph()
			path, dist := g.FordBelman(tt.start, tt.end)

			if tt.hasNegativeCycle && dist != -1 {
				if tt.expectedDist != -1 {
					t.Errorf("expected to detect negative cycle and return -1, got %d", dist)
				}
			} else {
				if dist != tt.expectedDist {
					t.Errorf("expected distance %d, got %d", tt.expectedDist, dist)
				}

				if !reflect.DeepEqual(path, tt.expectedPath) {
					t.Errorf("expected path %v, got %v", tt.expectedPath, path)
				}

				if len(path) > 0 {
					if path[0] != tt.start || path[len(path)-1] != tt.end {
						t.Errorf("path should start at %d and end at %d, got %v",
							tt.start, tt.end, path)
					}

					if dist != -1 {
						var sum int64
						current := g.GetVertex(path[0])
						for i := 1; i < len(path); i++ {
							next := g.GetVertex(path[i])
							found := false
							for _, edge := range current.Edges {
								if edge.To == next {
									sum += edge.Weight
									found = true
									break
								}
							}
							if !found {
								t.Errorf("missing edge %d→%d", current.Key, next.Key)
							}
							current = next
						}
						if dist != sum {
							t.Errorf("path weight sum %d != distance %d", sum, dist)
						}
					}
				}
			}
		})
	}
}
