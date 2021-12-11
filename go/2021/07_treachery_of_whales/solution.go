package day7

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
)

type TreacheryOfWhales struct{}

func (TreacheryOfWhales) Input() string {
	return "2021/07_treachery_of_whales/input.txt"
}

func (TreacheryOfWhales) SolveQ1(rl utils.ReadLine) string {
	arr := utils.OrderedIntArray{0}
	for {
		ch, eof := rl()
		if eof {
			break
		}
		number, _ := strconv.Atoi(ch)
		arr.Append(number)
	}

	fmt.Println(arr)
	middle := arr[utils.Middle(arr)]

	fuel := -middle
	for _, v := range arr {
		diff := utils.AbsInt(v - middle)
		fuel += int(diff)
	}
	return strconv.Itoa(fuel)
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
func readToEnd(rl utils.ReadLine) (arr []int, max int) {
	for {
		line, eof := rl()
		if eof {
			break
		}
		val, _ := strconv.Atoi(line)
		if val > max {
			max = val
		}
		arr = append(arr, val)
	}
	return arr, max
}
func (TreacheryOfWhales) SolveQ2(rl utils.ReadLine) string {
	arr, max := readToEnd(rl)
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
	return strconv.Itoa(min)
}
