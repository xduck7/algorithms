package dfs

import (
	"algotithms/graph/graph"
	"fmt"
)

type Graph struct {
	graph.Graph
}

func (g *Graph) DFS(startKey int) string {
	res := ""
	startVertex := g.GetVertex(startKey)
	if startVertex == nil {
		return "no vertexes"
	}

	visited := make(map[int]bool)
	stack := []*graph.Vertex{startVertex}

	for len(stack) > 0 {
		n := len(stack) - 1
		currentVertex := stack[n]
		stack = stack[:n]

		if !visited[currentVertex.Key] {
			res += fmt.Sprintf("%d", currentVertex.Key)
			visited[currentVertex.Key] = true

			for i := len(currentVertex.Near) - 1; i >= 0; i-- {
				near := currentVertex.Near[i]
				if !visited[near.Key] {
					stack = append(stack, near)
				}
			}
		}
	}
	return res
}
