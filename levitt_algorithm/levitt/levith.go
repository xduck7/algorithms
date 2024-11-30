package levitt

import "math"

type Levitt struct {
	Graph   *Graph
	Sources []int
	Dists   map[int]int
}

func NewLevitt(graph *Graph, sources []int) *Levitt {
	dists := make(map[int]int)
	for _, v := range graph.Vertices {
		dists[v] = math.MaxInt
	}
	for _, source := range sources {
		dists[source] = 0
	}

	return &Levitt{
		Graph:   graph,
		Sources: sources,
		Dists:   dists,
	}
}

func (l *Levitt) RelaxEdge(u, v int) {
	if l.Dists[v] > l.Dists[u]+l.Graph.Edges[u][v] {
		l.Dists[v] = l.Dists[u] + l.Graph.Edges[u][v]
	}
}

func (l *Levitt) Run() {
	for {
		updated := false
		for u := range l.Graph.Edges {
			for v, weight := range l.Graph.Edges[u] {
				if l.Dists[u] != math.MaxInt && l.Dists[v] > l.Dists[u]+weight {
					l.RelaxEdge(u, v)
					updated = true
				}
			}
		}
		if !updated {
			break
		}
	}
}

func (l *Levitt) GetDistances() map[int]int {
	return l.Dists
}
