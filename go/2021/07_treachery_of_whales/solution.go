package day7

import (
	"adventofcode/utils"
	"strconv"
)

type TreacheryOfWhales struct{}

func (TreacheryOfWhales) Path() string {
	return "2021/07_treachery_of_whales"
}

func (TreacheryOfWhales) SolveQ1(input []byte) int {
	read := utils.Reader(input, utils.CommaSplitFunc)
	arr := utils.OrderedIntArray{0}

	for ch, eof := read(); !eof; ch, eof = read() {
		number, _ := strconv.Atoi(ch)
		arr.Append(number)
	}

	middle := arr[utils.Middle(arr)]

	fuel := -middle
	for _, v := range arr {
		diff := utils.AbsInt(v - middle)
		fuel += int(diff)
	}
	return fuel
}

var cache = map[int]int{}

func calculateCost(target int, origin int) int {
	diff := utils.AbsInt(target - origin)
	if result, ok := cache[diff]; ok {
		return result
	}

	result := (diff * (diff + 1)) / 2
	cache[diff] = result
	return result
}
func readToEnd(input []byte) (arr []int, max int) {
	read := utils.Reader(input, utils.CommaSplitFunc)
	for ch, eof := read(); !eof; ch, eof = read() {
		val, _ := strconv.Atoi(ch)
		if val > max {
			max = val
		}
		arr = append(arr, val)
	}
	return arr, max
}
func (TreacheryOfWhales) SolveQ2(input []byte) int {
	arr, max := readToEnd(input)
	min := 99999999
	for index := 0; index < max; index++ {
		cost := 0
		for _, v := range arr {
			cost += calculateCost(index, v)
		}
		if cost < min {
			min = cost
		}
	}
	return min
}
