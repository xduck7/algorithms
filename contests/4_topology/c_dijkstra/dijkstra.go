package dijkstra

import (
	"container/heap"
	"math"
)

func (g *Graph) Dijkstra(start, end int64) (int64, []int64) {
	dist := make(map[*Vertex]int64)
	prev := make(map[*Vertex]*Vertex)
	q := make(Queue, 0)
	heap.Init(&q)

	startVertex := g.GetVertex(start)
	endVertex := g.GetVertex(end)

	if startVertex == nil || endVertex == nil {
		return -1, nil
	}

	for _, v := range g.Vertices {
		dist[v] = math.MaxInt64
	}
	dist[startVertex] = 0

	heap.Push(&q, &QueueItem{
		vertex:   startVertex,
		priority: 0,
	})

	for q.Len() > 0 {
		current := heap.Pop(&q).(*QueueItem).vertex

		if current == endVertex {
			break
		}

		for _, edge := range current.Edges {
			neighbor := edge.To
			newDist := dist[current] + edge.Weight

			if newDist < dist[neighbor] {
				dist[neighbor] = newDist
				prev[neighbor] = current
				heap.Push(&q, &QueueItem{
					vertex:   neighbor,
					priority: int(newDist),
				})
			}
		}
	}

	if dist[endVertex] == math.MaxInt64 {
		return -1, nil
	}

	path := []int64{}
	for v := endVertex; v != nil; v = prev[v] {
		path = append(path, v.Key)
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return dist[endVertex], path
}
