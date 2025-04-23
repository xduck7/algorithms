package bfs

import (
	"algotithms/graph/graph"
	"testing"
)

type MockGraph struct {
	Vertices map[int]*graph.Vertex
}

func NewMockGraph() *MockGraph {
	return &MockGraph{Vertices: make(map[int]*graph.Vertex)}
}

func (g *MockGraph) AddVertex(key int) {
	if _, exists := g.Vertices[key]; !exists {
		g.Vertices[key] = &graph.Vertex{Key: key}
	}
}

func (g *MockGraph) AddEdge(from, to int) {
	fromVertex, fromExists := g.Vertices[from]
	toVertex, toExists := g.Vertices[to]

	if fromExists && toExists {
		fromVertex.Adjacent = append(fromVertex.Adjacent, toVertex)
	}
}

func (g *MockGraph) GetVertex(key int) *graph.Vertex {
	return g.Vertices[key]
}

func (g *MockGraph) ContainsVertex(key int) bool {
	for _, v := range g.Vertices {
		if v.Key == key {
			return true
		}
	}
	return false
}

func TestBFS(t *testing.T) {
	// Создаем граф
	graph := NewMockGraph()
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)
	graph.AddVertex(5)

	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 5)

	bfsGraph := &Graph{graph}
	res := bfsGraph.BFS(1)

	expected := "Visited 1\nVisited 2\nVisited 3\nVisited 4\nVisited 5\n"
	if res != expected {
		t.Errorf("expected %q, but got %q", expected, res)
	}
}
