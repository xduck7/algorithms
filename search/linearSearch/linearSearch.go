package linearSearch

func linearSearch(sortedList []int, target int) bool {

	for i := range sortedList {
		if sortedList[i] == target {
			return true
		}
	}
	return false
}
