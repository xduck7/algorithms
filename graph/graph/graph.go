package graph

type Vertex struct {
	Key  int
	Near []*Vertex
}

type Graph struct {
	Vertices []*Vertex
}

func (g *Graph) AddVertex(key int) {
	if !g.ContainsVertex(key) {
		g.Vertices = append(g.Vertices, &Vertex{Key: key})
	}
}

func (g *Graph) AddEdge(from, to int) {
	fromVertex := g.GetVertex(from)
	toVertex := g.GetVertex(to)

	if fromVertex != nil && toVertex != nil {
		fromVertex.Near = append(fromVertex.Near, toVertex)
	}
}

func (g *Graph) GetVertex(key int) *Vertex {
	for _, v := range g.Vertices {
		if v.Key == key {
			return v
		}
	}
	return nil
}

func (g *Graph) ContainsVertex(key int) bool {
	for _, v := range g.Vertices {
		if v.Key == key {
			return true
		}
	}
	return false
}
