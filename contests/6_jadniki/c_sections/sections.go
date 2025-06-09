package sections

import (
	"sort"
	"strconv"
	"strings"
)

func MaxSectionStrick(n int, sec string) int {
	lines := strings.Split(strings.TrimSpace(sec), "\n")
	linesList := make([][2]int, n)
	for i := 0; i < n; i++ {
		otrezok := strings.Fields(lines[i])
		l, _ := strconv.Atoi(otrezok[0])
		r, _ := strconv.Atoi(otrezok[1])
		linesList[i][0] = l
		linesList[i][1] = r
	}

	sort.Slice(linesList, func(i, j int) bool {
		return linesList[i][1] < linesList[j][1]
	})

	count := 0
	end := -1
	for i := 0; i < n; i++ {
		if linesList[i][0] >= end {
			count++
			end = linesList[i][1]
		}
	}

	return count
}
