package bfs_test

import (
	"algotithms/contests/3_dfs_bfs/bfs"
	"algotithms/graph/graph"
	"testing"
)

func TestBFS(t *testing.T) {
	g := graph.Graph{}
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)

	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)

	bfsGraph := &bfs.Graph{g}
	res := bfsGraph.BFS(1)

	expected := "12345"
	if res != expected {
		t.Errorf("expected: %s but got: %s", expected, res)
	}
}
