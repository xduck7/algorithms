package exponentialSearch

func exponentialSearch(sortedArray []int, target int) bool {
	length := len(sortedArray)
	if length == 0 {
		return false
	}
	if sortedArray[0] == target {
		return true
	}
	i := 1
	for i < length && sortedArray[i] <= target {
		i = i * 2
	}
	return binarySearch(sortedArray, target, i/2, min(i, length-1))
}

func binarySearch(sortedArray []int, target, l, r int) bool {

	left := l
	right := r

	if r >= l {
		mid := left + (right-left)/2

		if sortedArray[mid] == target {
			return true
		}
		if sortedArray[mid] > target {
			return binarySearch(sortedArray, target, left, mid-1)
		}
		return binarySearch(sortedArray, target, mid+1, right)
	}

	return false
}
