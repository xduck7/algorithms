package levitt

type Graph struct {
	Vertices []int
	Edges    map[int]map[int]int
}

func NewGraph(vertices []int) *Graph {
	edges := make(map[int]map[int]int)
	for _, v := range vertices {
		edges[v] = make(map[int]int)
	}
	return &Graph{
		Vertices: vertices,
		Edges:    edges,
	}
}

func (g *Graph) AddEdge(from, to, weight int) {
	g.Edges[from][to] = weight
}
