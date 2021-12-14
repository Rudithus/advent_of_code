package day1

import (
	"adventofcode/utils"
	"bufio"
	"strconv"
)

type Sonarsweep struct{}

func (Sonarsweep) Path() string {
	return "2021/01_sonar_sweep"
}

func (Sonarsweep) SolveQ1(input []byte) int {
	counter := 0
	previous := 9999
	read := utils.Reader(input, bufio.ScanLines)

	for line, eof := read(); !eof; line, eof = read() {
		depth, _ := strconv.Atoi(line)
		if depth > previous {
			counter++
		}
		previous = depth
	}

	return counter
}

func parseDepths(input []byte) []int {
	var arr []int
	read := utils.Reader(input, bufio.ScanLines)

	for line, eof := read(); !eof; line, eof = read() {
		val, err := strconv.Atoi(line)
		utils.Check(err)

		arr = append(arr, val)
	}

	return arr
}

func (Sonarsweep) SolveQ2(input []byte) int {
	depths := parseDepths(input)

	counter := 0
	previous := utils.SumInts(depths[0:3])

	for i := 4; i <= len(depths); i++ {
		sum := utils.SumInts(depths[i-3 : i])
		if sum > previous {
			counter++
		}
		previous = sum
	}

	return counter
}
