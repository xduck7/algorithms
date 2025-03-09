package peak

func findPeak(arr []int64) int64 {
	if len(arr) == 0 {
		return -1
	}
	left := 0
	right := len(arr) - 1

	for left < right {
		mid := left + (right-left)/2
		if arr[mid] > arr[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return int64(left)
}
