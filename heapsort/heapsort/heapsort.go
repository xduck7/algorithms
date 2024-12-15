package heapsort

func heapify(arr []int, n, index int) {
	largest := index
	left := 2*index + 1
	right := 2*index + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != index {
		arr[index], arr[largest] = arr[largest], arr[index]
		heapify(arr, n, largest)
	}
}

func Heapsort(arr []int) {
	n := len(arr)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]

		heapify(arr, i, 0)
	}
}
