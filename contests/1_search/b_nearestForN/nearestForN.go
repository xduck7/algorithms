package nearestForX

import "fmt"

func checkArea(val, radius, target int64) bool {
	return (val-radius == target) || (val+radius == target)
}

func nearestForX(arr []int64, x int64) int64 {
	for radius := 0; radius < 1000; radius++ {
		for index := range arr {
			if checkArea(arr[index], int64(radius), x) {
				fmt.Printf("R: %d", radius)
				return int64(index)
			}
		}
	}
	return -1
}
