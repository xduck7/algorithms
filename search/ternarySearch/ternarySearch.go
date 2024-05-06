package ternarySearch

func ternarySearch(sortedArray []int, target int) bool {
	start := 0
	end := len(sortedArray) - 1
	for start <= end {
		middle := (end - start) / 3
		left := start + middle
		right := end - middle

		leftVal := sortedArray[left]
		rightVal := sortedArray[right]

		if leftVal == target {
			return true
		}
		if rightVal == target {
			return true
		}
		if leftVal < target && target < rightVal {
			start = left + 1
			end = right - 1
		}
		if target < leftVal {
			end = right - 1
		} else {
			start = left + 1
		}
	}
	return false
}
