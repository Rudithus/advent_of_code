package utils

import "math"

func Middle(arr []int) int {
	return int(math.Ceil(float64(len(arr)) / 2))
}

func Sum(arr []int) int {
	sum := 0
	for _, i := range arr {
		sum += i
	}
	return sum
}
