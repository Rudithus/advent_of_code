package day1

import (
	"fmt"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func sum(arr []int) int {
	sum := 0
	for _, i := range arr {
		sum += i
	}
	return sum
}

func convert(lines []string) []int {
	var numbers []int
	for _, l := range lines {
		i, err := strconv.Atoi(l)
		check(err)
		numbers = append(numbers, i)
	}
	return numbers
}

func Solve2(s string) int {
	lines := strings.Fields(s)
	depths := convert(lines)

	counter := 0
	previous := sum(depths[0:3])

	for i := 4; i <= len(depths); i++ {
		sum := sum(depths[i-3 : i])
		if sum > previous {
			counter++
		}
		previous = sum
	}

	fmt.Println(counter)
	return counter
}
