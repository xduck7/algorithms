package quickSort

func quickSort(array []int) []int {
	if len(array) <= 1 {
		return nil
	}

	mid := sort(array)

	quickSort(array[:mid])
	quickSort(array[mid+1:])
	return array
}

func sort(array []int) int {
	pivot := array[len(array)-1]
	i := -1

	for j := 0; j < len(array)-1; j++ {
		if array[j] < pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}

	array[i+1], array[len(array)-1] = array[len(array)-1], array[i+1]
	return i + 1
}
