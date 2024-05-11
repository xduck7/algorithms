package heapSort

func heapSort(array []int) []int {
	length := len(array)

	for i := length/2 - 1; i >= 0; i-- {
		heap(array, length, i)
	}

	for i := length - 1; i > 0; i-- {
		array[0], array[i] = array[i], array[0]
		heap(array, i, 0)
	}

	return array
}

func heap(array []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && array[left] > array[largest] {
		largest = left
	}

	if right < n && array[right] > array[largest] {
		largest = right
	}

	if largest != i {
		array[i], array[largest] = array[largest], array[i]

		heap(array, n, largest)
	}
}
