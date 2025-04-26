package shortest_path

func ShortestPath(n, a, b int, g [][]int) (int, []int) {
	d := make([]int, n+1)
	p := make([]int, n+1)
	for i := range d {
		d[i] = -1
	}
	d[a] = 0

	q := []int{a}
	founded := false

	for len(q) > 0 && !founded {
		u := q[0]
		q = q[1:]

		for _, v := range g[u] {
			if d[v] == -1 {
				d[v] = d[u] + 1
				p[v] = u
				q = append(q, v)

				if v == b {
					founded = true
					break
				}
			}
		}
	}

	if d[b] == -1 {
		return -1, nil
	}

	path := []int{}
	for v := b; v != 0; v = p[v] {
		path = append(path, v)
	}

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return d[b], path
}
