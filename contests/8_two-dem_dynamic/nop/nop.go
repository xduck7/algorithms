package nop

func Nop(a, b string) string {
	n := len(a)
	m := len(b)

	mn := make([][]int, n+1)
	for i := range mn {
		mn[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i-1] == b[j-1] {
				mn[i][j] = mn[i-1][j-1] + 1
			} else if mn[i-1][j] > mn[i][j-1] {
				mn[i][j] = mn[i-1][j]
			} else {
				mn[i][j] = mn[i][j-1]
			}
		}
	}

	i, j := n, m
	res := make([]byte, 0)
	for i > 0 && j > 0 {
		if a[i-1] == b[j-1] {
			res = append([]byte{a[i-1]}, res...)
			i--
			j--
		} else if mn[i-1][j] > mn[i][j-1] {
			i--
		} else {
			j--
		}
	}

	return string(res)
}
