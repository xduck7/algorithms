package maxitems

import (
	"sort"
	"strconv"
	"strings"
)

func MaxItems(nk, prices string) int {
	chars := strings.Fields(nk)
	n, _ := strconv.Atoi(chars[0])
	k, _ := strconv.Atoi(chars[1])

	pricesString := strings.Fields(prices)
	pricesList := make([]int, n)
	for i := 0; i < n; i++ {
		pricesList[i], _ = strconv.Atoi(pricesString[i])
	}

	sort.Ints(pricesList)

	count := 0
	for _, cost := range pricesList {
		if k >= cost {
			k -= cost
			count++
		} else {
			break
		}
	}

	return count
}
