package prim

type Edge struct {
	From   int64
	To     int64
	Weight int64
}

type Graph struct {
	Vertices []int64
	Edges    []Edge
}

func (g *Graph) AddVertex(v int64) {
	g.Vertices = append(g.Vertices, v)
}

func (g *Graph) AddEdge(u, v, w int64) {
	g.Edges = append(g.Edges, Edge{u, v, w})
	g.Edges = append(g.Edges, Edge{v, u, w})
}

func (g *Graph) GetAdjacencyList() map[int64][]Edge {
	adj := make(map[int64][]Edge)
	for _, e := range g.Edges {
		adj[e.From] = append(adj[e.From], e)
	}
	return adj
}
