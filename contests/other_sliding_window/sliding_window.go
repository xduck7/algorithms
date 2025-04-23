package other_sliding_window

func findLenOfUniqSubstring(s []rune) int64 {
	if len(s) == 0 {
		return 0
	}
	windowContent := make(map[rune]int)
	left := 0
	maxLen := 0

	for right, letter := range s {
		if index, ok := windowContent[letter]; ok && index >= left {
			left = index + 1
		}
		windowContent[letter] = right
		currentLen := right - left + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return int64(maxLen)
}
