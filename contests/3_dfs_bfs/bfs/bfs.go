package bfs

import (
	"algotithms/graph/graph"
	"container/list"
	"fmt"
)

type Graph struct {
	graph.Graph
}

func (g *Graph) BFS(startKey int) string {
	res := ""
	startVertex := g.GetVertex(startKey)
	if startVertex == nil {
		return "No vertex found"
	}

	visited := make(map[int]bool)
	queue := list.New()
	queue.PushBack(startVertex)

	for queue.Len() > 0 {
		element := queue.Front()
		currentVertex := element.Value.(*graph.Vertex)
		queue.Remove(element)

		if !visited[currentVertex.Key] {
			res += fmt.Sprintf("%d", currentVertex.Key)
			visited[currentVertex.Key] = true

			for _, v := range currentVertex.Near {
				if !visited[v.Key] {
					queue.PushBack(v)
				}
			}
		}
	}
	return res
}
