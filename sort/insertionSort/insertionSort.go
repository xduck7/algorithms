package insertionSort

func insertionSort(array []int) []int {
	for left := 1; left < len(array); left++ {
		key := array[left]
		right := left - 1
		for right >= 0 && array[right] > key {
			array[right+1] = array[right]
			right--
		}
		array[right+1] = key
	}
	return array
}
