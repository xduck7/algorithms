package binarySearch

func binarySearch(sortedList []int, target int) bool {

	left := 0
	right := len(sortedList) - 1

	for left <= right {

		mid := (left + right) / 2
		midElement := sortedList[mid]

		if midElement == target {
			return true
		}
		if target > midElement {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
