package kana

func Kana(N int, v [][]int) []int {
	in := make([]int, N+1)
	for u := 1; u <= N; u++ {
		for _, val := range v[u] {
			in[val]++
		}
	}

	var queue []int
	for u := 1; u <= N; u++ {
		if in[u] == 0 {
			queue = append(queue, u)
		}
	}

	var res []int
	c := 0

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		res = append(res, u)
		c++

		for _, val := range v[u] {
			in[val]--
			if in[val] == 0 {
				queue = append(queue, val)
			}
		}
	}

	if c != N {
		return nil
	}

	return res
}
