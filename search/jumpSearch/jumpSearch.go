package jumpSearch

import (
	"math"
)

func jumpSearch(sortedArray []int, target int) bool {
	step := int(math.Sqrt(float64(len(sortedArray))))
	length := len(sortedArray)
	left := 0
	right := step
	for right < length && sortedArray[right-1] < target {
		left = right
		right += step
	}
	for i := left; i < right-1; i++ {
		if sortedArray[i] == target {
			return true
		}
	}
	return false
}
