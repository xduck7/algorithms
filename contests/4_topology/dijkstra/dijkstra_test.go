package dijkstra_test

import (
	"algotithms/contests/4_topology/dijkstra"
	"fmt"
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	tests := []struct {
		graph        func() *dijkstra.Graph
		start        int64
		end          int64
		expectedDist int64
		expectedPath []int64
	}{
		{
			graph: func() *dijkstra.Graph {
				g := &dijkstra.Graph{}
				g.AddVertex(1)
				g.AddVertex(2)
				g.AddVertex(3)
				g.AddWeightedEdge(1, 2, 5)
				g.AddWeightedEdge(1, 3, 3)
				g.AddWeightedEdge(2, 3, 1)
				return g
			},
			start:        1,
			end:          3,
			expectedDist: 3,
			expectedPath: []int64{1, 3},
		},
		{
			graph: func() *dijkstra.Graph {
				g := &dijkstra.Graph{}
				g.AddVertex(1)
				g.AddVertex(2)
				g.AddVertex(3)
				g.AddWeightedEdge(1, 2, 5)
				return g
			},
			start:        1,
			end:          3,
			expectedDist: -1,
			expectedPath: nil,
		},
		{
			graph: func() *dijkstra.Graph {
				g := &dijkstra.Graph{}
				g.AddVertex(1)
				g.AddVertex(2)
				g.AddVertex(3)
				g.AddVertex(4)
				g.AddWeightedEdge(1, 2, 1)
				g.AddWeightedEdge(1, 3, 4)
				g.AddWeightedEdge(2, 4, 2)
				g.AddWeightedEdge(3, 4, 1)
				return g
			},
			start:        1,
			end:          4,
			expectedDist: 3,
			expectedPath: []int64{1, 2, 4},
		},
		{
			graph: func() *dijkstra.Graph {
				g := &dijkstra.Graph{}
				g.AddVertex(1)
				g.AddVertex(2)
				g.AddWeightedEdge(1, 2, 5)
				return g
			},
			start:        1,
			end:          1,
			expectedDist: 0,
			expectedPath: []int64{1},
		},
		{
			graph: func() *dijkstra.Graph {
				g := &dijkstra.Graph{}
				for i := 1; i <= 5; i++ {
					g.AddVertex(int64(i))
				}
				g.AddWeightedEdge(1, 2, 2)
				g.AddWeightedEdge(1, 3, 4)
				g.AddWeightedEdge(2, 3, 1)
				g.AddWeightedEdge(2, 4, 7)
				g.AddWeightedEdge(3, 5, 3)
				g.AddWeightedEdge(5, 4, 2)
				return g
			},
			start:        1,
			end:          4,
			expectedDist: 8,
			expectedPath: []int64{1, 2, 3, 5, 4},
		},
	}

	for id, tt := range tests {
		t.Run(fmt.Sprintf("%d", id), func(t *testing.T) {
			g := tt.graph()
			dist, path := g.Dijkstra(tt.start, tt.end)

			if dist != tt.expectedDist {
				t.Errorf("expected distance %d, got %d", tt.expectedDist, dist)
			}

			if !reflect.DeepEqual(path, tt.expectedPath) {
				t.Errorf("expected path %v, got %v", tt.expectedPath, path)
			}

			if len(path) > 0 {
				if path[0] != tt.start || path[len(path)-1] != tt.end {
					t.Errorf("expected %d→%d, got %d→%d",
						tt.start, tt.end, path[0], path[len(path)-1])
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
							t.Errorf("edge %d→%d", current.Key, next.Key)
						}
						current = next
					}
					if dist != sum {
						t.Errorf("path weight sum %d =! %d", sum, dist)
					}
				}
			}
		})
	}
}
