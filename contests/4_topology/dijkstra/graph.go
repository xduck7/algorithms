package dijkstra

type Vertex struct {
	Key    int64
	Weight int64
	Near   []*Vertex
	Edges  []Edge
}

type Edge struct {
	To     *Vertex
	Weight int64
}

type Graph struct {
	Vertices []*Vertex
}

func (g *Graph) AddVertex(key int64) {
	if !g.ContainsVertex(key) {
		g.Vertices = append(g.Vertices, &Vertex{Key: key})
	}
}

func (g *Graph) AddEdge(from, to int64) {
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)

	if fromVertex != nil && toVertex != nil {
		fromVertex.Near = append(fromVertex.Near, toVertex)
	}
}

func (g *Graph) AddWeightedEdge(from, to, weight int64) {
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)

	if fromVertex != nil && toVertex != nil {
		fromVertex.Edges = append(fromVertex.Edges, Edge{
			To:     toVertex,
			Weight: weight,
		})
	}
}

func (g *Graph) GetVertex(key int64) *Vertex {
	for _, v := range g.Vertices {
		if v.Key == key {
			return v
		}
	}
	return nil
}

func (g *Graph) ContainsVertex(key int64) bool {
	for _, v := range g.Vertices {
		if v.Key == key {
			return true
		}
	}
	return false
}
