package longestprogression

import (
	"sort"
	"strconv"
	"strings"
)

func LongestProgression(n int, str string) int {
	nums := make([]int, n)
	fs := strings.Fields(str)
	for i := range n {
		nums[i], _ = strconv.Atoi(fs[i])
	}

	temp := []int{}
	for _, num := range nums {
		i := sort.SearchInts(temp, num)
		if i == len(temp) {
			temp = append(temp, num)
		} else {
			temp[i] = num
		}
	}
	return len(temp)
}
