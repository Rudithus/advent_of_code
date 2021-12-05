package day1

import (
	"adventofcode/utils"
	"strconv"
)

type Sonarsweep struct{}

func (Sonarsweep) Input() string {
	return "2021/01_sonar_sweep/input.txt"
}

func (Sonarsweep) SolveQ1(rl utils.ReadLine) string {
	counter := 0
	previous := 9999

	for {
		line, eof := rl()
		if eof {
			break
		}

		depth, _ := strconv.Atoi(line)
		if depth > previous {
			counter++
		}
		previous = depth
	}

	return strconv.Itoa(counter)
}

func readtoend(rl utils.ReadLine) []int {
	var arr []int
	for {
		line, eof := rl()
		if eof {
			break
		}

		val, _ := strconv.Atoi(line)
		arr = append(arr, val)
	}

	return arr
}

func (Sonarsweep) SolveQ2(rl utils.ReadLine) string {
	depths := readtoend(rl)

	counter := 0
	previous := utils.Sum(depths[0:3])

	for i := 4; i <= len(depths); i++ {
		sum := utils.Sum(depths[i-3 : i])
		if sum > previous {
			counter++
		}
		previous = sum
	}

	return strconv.Itoa(counter)
}
