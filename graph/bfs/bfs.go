package bfs

import (
	"algotithms/graph/graph"
	"container/list"
	"fmt"
	"io"
)

type Graph struct {
	graph.IGraph
}

func (g *Graph) BFS(startKey int, writer io.Writer) {
	startVertex := g.GetVertex(startKey)
	if startVertex == nil {
		fmt.Fprintln(writer, "Start vertex not found")
		return
	}

	visited := make(map[int]bool)
	queue := list.New()
	queue.PushBack(startVertex)

	for queue.Len() > 0 {
		element := queue.Front()
		currentVertex := element.Value.(*graph.Vertex)
		queue.Remove(element)

		if !visited[currentVertex.Key] {
			fmt.Fprintf(writer, "Visited %d\n", currentVertex.Key)
			visited[currentVertex.Key] = true

			for _, v := range currentVertex.Adjacent {
				if !visited[v.Key] {
					queue.PushBack(v)
				}
			}
		}
	}
}
