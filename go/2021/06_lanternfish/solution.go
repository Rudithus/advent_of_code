package day6

import (
	"adventofcode/utils"
	"bufio"
	"strconv"
	"strings"
)

type Lanternfish struct{}

func (Lanternfish) Path() string {
	return "2021/06_lanternfish"
}

func parseFish(line string) [9]int {
	var fishes [9]int
	for _, v := range strings.Split(line, ",") {
		n, _ := strconv.Atoi(v)
		fishes[n]++
	}
	return fishes
}

func iterate(fishes []int) {
	newFish := fishes[0]

	for i, fishCount := range fishes[1:] {
		fishes[i] = fishCount
	}
	fishes[8] = 0
	fishes[6] += newFish
	fishes[8] += newFish
}

func (Lanternfish) SolveQ1(input []byte) int {
	read := utils.Reader(input, bufio.ScanLines)
	line, _ := read()

	fishes := parseFish(line)

	for i := 0; i < 80; i++ {
		iterate(fishes[:])
	}

	return utils.SumInts(fishes[:])
}

func (Lanternfish) SolveQ2(input []byte) int {
	read := utils.Reader(input, bufio.ScanLines)
	line, _ := read()

	fishes := parseFish(line)

	for i := 0; i < 256; i++ {
		iterate(fishes[:])
	}

	return utils.SumInts(fishes[:])
}
