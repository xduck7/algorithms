package interpolationSearch

func interpolationSearch(sortedArray []int, target int) bool {
	left := 0
	right := len(sortedArray) - 1
	var mid int
	for left <= right && target >= sortedArray[left] && target <= sortedArray[right] {
		mid = left + ((right-left)*(target-sortedArray[left]))/(sortedArray[right]-sortedArray[left])

		if sortedArray[mid] < target {
			left = mid + 1
		} else if sortedArray[mid] > target {
			right = mid - 1
		} else {
			return target == sortedArray[mid]
		}
	}
	return false
}
