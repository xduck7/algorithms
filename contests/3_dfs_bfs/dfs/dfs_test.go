package dfs_test

import (
	"algotithms/contests/3_dfs_bfs/dfs"
	"algotithms/graph/graph"
	"testing"
)

func TestDFS(t *testing.T) {
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

	dfsGraph := dfs.Graph{g}
	res := dfsGraph.DFS(1)

	expected := []string{
		"12453",
		"13452",
	}

	match := false
	for _, v := range expected {
		if res == v {
			match = true
			break
		}
	}

	if !match {
		t.Errorf("expected one of %v, but got %q", expected, res)
	}
}
