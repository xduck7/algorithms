package fibonacciSearch

func fibonacciSearch(sortedArray []int, target int) bool {
	length := len(sortedArray)
	if length == 0 {
		return false
	}

	f0 := 0
	f1 := 1
	f2 := f0 + f1

	for f2 < length {
		f0 = f1
		f1 = f2
		f2 = f0 + f1
	}

	offset := -1

	for f1 > 1 {
		i := min(offset+f0, length-1)
		if sortedArray[i] < target {
			f2 = f1
			f1 = f0
			f0 = f2 - f1
			offset = i
		} else if sortedArray[i] > target {
			f2 = f0
			f1 = f1 - f0
			f0 = f2 - f1
		} else {
			return true
		}
	}

	if f1 == 1 && sortedArray[offset+1] == target {
		return true
	}

	return false
}
