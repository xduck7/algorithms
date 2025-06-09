package ford_belman

import "math"

func (g *Graph) FordBelman(startKey, endKey int64) ([]int64, int64) {
	dist := make(map[int64]int64)
	p := make(map[int64]int64)

	for _, v := range g.Vertices {
		dist[v.Key] = math.MaxInt64
	}
	dist[startKey] = 0

	for i := 0; i < len(g.Vertices)-1; i++ {
		for _, u := range g.Vertices {
			for _, edge := range u.Edges {
				v := edge.To
				if dist[u.Key] != math.MaxInt64 && dist[v.Key] > dist[u.Key]+edge.Weight {
					dist[v.Key] = dist[u.Key] + edge.Weight
					p[v.Key] = u.Key
				}
			}
		}
	}

	for _, val := range g.Vertices {
		for _, edge := range val.Edges {
			v := edge.To
			if dist[val.Key] != math.MaxInt64 && dist[v.Key] > dist[val.Key]+edge.Weight {
				return nil, -1
			}
		}
	}

	if dist[endKey] == math.MaxInt64 {
		return nil, -1
	}

	var path []int64
	current := endKey
	for current != startKey {
		path = append([]int64{current}, path...)
		current = p[current]
	}
	path = append([]int64{startKey}, path...)

	return path, dist[endKey]
}
