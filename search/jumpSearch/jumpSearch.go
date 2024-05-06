package jumpSearch

import (
	"math"
)

func jumpSearch(sortedList []int, target int) bool {
	step := int(math.Sqrt(float64(len(sortedList))))
	length := len(sortedList)
	start := 0
	end := step
	for end < length && sortedList[end-1] < target {
		start = end
		end += step
	}
	for i := start; i < end-1; i++ {
		if sortedList[i] == target {
			return true
		}
	}
	return false
}
