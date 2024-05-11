package bubbleSort

func bubbleSort(array []int) []int {
	length := len(array)
	for i := 0; i < length-1; {
		for j := 0; j < length-1-i; j++ {
			if array[j] > array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
		i++
	}
	return array
}
