package first

func firstNotLowerN(arr []int, x int) int {
	low, high := 0, len(arr)-1
	result := -1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid] >= x {
			result = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return result
}
