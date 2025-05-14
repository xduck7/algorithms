package prim

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrim(t *testing.T) {
	tests := []struct {
		graphInit         func() *Graph
		expectedWeight    int64
		expectedEdges     []Edge
		expectedConnected bool
	}{
		{
			graphInit: func() *Graph {
				g := &Graph{}
				g.AddVertex(1)
				g.AddVertex(2)
				g.AddVertex(3)
				g.AddEdge(1, 2, 1)
				g.AddEdge(2, 3, 2)
				g.AddEdge(1, 3, 3)
				return g
			},
			expectedWeight: 3,
			expectedEdges: []Edge{
				{2, 3, 2},
				{1, 2, 1},
			},
			expectedConnected: true,
		},
		{
			graphInit: func() *Graph {
				g := &Graph{}
				g.AddVertex(1)
				g.AddVertex(2)
				g.AddVertex(3)
				g.AddEdge(1, 2, 1)
				return g
			},
			expectedWeight:    -1,
			expectedEdges:     nil,
			expectedConnected: false,
		},
		{
			graphInit: func() *Graph {
				g := &Graph{}
				g.AddVertex(1)
				return g
			},
			expectedWeight:    0,
			expectedEdges:     []Edge{},
			expectedConnected: true,
		},
		{
			graphInit: func() *Graph {
				g := &Graph{}
				for i := 1; i <= 5; i++ {
					g.AddVertex(int64(i))
				}
				g.AddEdge(1, 2, 2)
				g.AddEdge(1, 3, 4)
				g.AddEdge(2, 3, 1)
				g.AddEdge(2, 4, 7)
				g.AddEdge(3, 5, 3)
				g.AddEdge(5, 4, 2)
				return g
			},
			expectedWeight: 8,
			expectedEdges: []Edge{
				{5, 4, 2},
				{3, 5, 3},
				{2, 3, 1},
				{1, 2, 2},
			},
			expectedConnected: true,
		},
	}

	for id, tt := range tests {
		t.Run(fmt.Sprintf("%d", id), func(t *testing.T) {
			g := tt.graphInit()
			weight, edges, connected := g.Prim()

			if connected != tt.expectedConnected {
				t.Errorf("expected connected %v, got %v", tt.expectedConnected, connected)
			}

			if weight != tt.expectedWeight {
				t.Errorf("expected weight %d, got %d", tt.expectedWeight, weight)
			}

			if !reflect.DeepEqual(edges, tt.expectedEdges) {
				t.Errorf("expected edges %v, got %v", tt.expectedEdges, edges)
			}

			if connected && weight != -1 {
				var sum int64
				for _, e := range edges {
					sum += e.Weight
				}
				if sum != weight {
					t.Errorf("sum of edge weights (%d) != total weight (%d)", sum, weight)
				}
			}
		})
	}
}
