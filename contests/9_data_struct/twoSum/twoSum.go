package twoSum

func TwoSum(a []int, k int) bool {
	s := make(map[int]bool)

	for _, num := range a {
		if s[k-num] {
			return true
		}
		s[num] = true
	}

	return false
}
