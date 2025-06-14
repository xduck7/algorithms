package containers

func Containers(n, k int64, w []int64) int64 {
	maxW := int64(0)
	sum := int64(0)
	for _, x := range w {
		if x > maxW {
			maxW = x
		}
		sum += x
	}

	l, r := maxW, sum
	for l < r {
		m := (l + r) / 2
		if ok(w, k, m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func ok(w []int64, k, cap int64) bool {
	t, curr := int64(1), int64(0)
	for _, x := range w {
		if curr+x <= cap {
			curr += x
		} else {
			t++
			curr = x
			if t > k {
				return false
			}
		}
	}
	return true
}
