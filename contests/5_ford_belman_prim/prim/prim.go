package prim

import (
	"container/heap"
	"sort"
)

func (g *Graph) Prim() (int64, []Edge, bool) {
	if len(g.Vertices) == 0 {
		return 0, nil, true
	}

	visited := make(map[int64]bool)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	start := g.Vertices[0]
	heap.Push(&pq, &Item{
		Node:     start,
		Priority: 0,
		Parent:   -1},
	)

	totalWeight := int64(0)
	mstEdges := make([]Edge, 0)
	adj := g.GetAdjacencyList()

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		node := item.Node

		if visited[node] {
			continue
		}

		visited[node] = true
		totalWeight += item.Priority

		if item.Parent != -1 {
			mstEdges = append(mstEdges, Edge{
				item.Parent,
				node,
				item.Priority})
		}

		for _, e := range adj[node] {
			if !visited[e.To] {
				heap.Push(&pq, &Item{
					Node:     e.To,
					Priority: e.Weight,
					Parent:   node,
				})
			}
		}
	}

	if len(visited) != len(g.Vertices) {
		return -1, nil, false
	}

	sortEdges(mstEdges)

	return totalWeight, mstEdges, true
}

func sortEdges(edges []Edge) {
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].From == edges[j].From {
			return edges[i].To > edges[j].To
		}
		return edges[i].From > edges[j].From
	})
}
