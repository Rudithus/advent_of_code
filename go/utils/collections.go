package utils

import (
	"fmt"
	"math"
)

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

func Contains(arr *[]string, value string) bool {
	for _, v := range *arr {
		if v == value {
			return true
		}
	}
	return false
}

func PrintGrid(grid [][]bool) {
	for _, arr := range grid {
		str := ""
		for _, dot := range arr {
			if dot {
				str += "#"
			} else {
				str += "."
			}
		}
		fmt.Println(str)
	}
}
