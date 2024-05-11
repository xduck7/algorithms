package selectionSort

func selectionSort(array []int) []int {
	length := len(array)
	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			array[i], array[minIndex] = array[minIndex], array[i]
		}
	}
	return array
}
