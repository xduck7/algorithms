package binarySearch

func binarySearch(sortedArray []int, target int) bool {

	left := 0
	right := len(sortedArray) - 1

	for left <= right {

		mid := (left + right) / 2
		midElement := sortedArray[mid]

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
