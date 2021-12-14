package utils

import "math"

func Middle(arr []int) int {
	return int(math.Ceil(float64(len(arr)) / 2))
}

func SumInts(arr []int) int {
	sum := 0
	for _, i := range arr {
		sum += i
	}
	return sum
}
func MultiplyBytes(arr []byte) int {
	multiply := 1
	for _, i := range arr {
		multiply *= int(i)
	}
	return multiply
}
