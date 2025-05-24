package maxway

import (
	"strconv"
	"strings"
)

func MaxWay(n int, str string) int {
	elems := strings.Fields(str)
	norm := make([]int, n)
	for i := range norm {
		norm[i], _ = strconv.Atoi(elems[i])
	}

	if n == 1 {
		return norm[0]
	}

	m := make([]int, n)
	m[0] = norm[0]
	m[1] = max(norm[0]+norm[1], norm[1])

	for i := 2; i < n; i++ {
		m[i] = max(m[i-1]+norm[i], m[i-2]+norm[i])
	}

	return m[n-1]
}
