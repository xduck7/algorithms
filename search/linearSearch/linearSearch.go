package linearSearch

func linearSearch(sortedArray []int, target int) bool {

	for i := range sortedArray {
		if sortedArray[i] == target {
			return true
		}
	}
	return false
}
