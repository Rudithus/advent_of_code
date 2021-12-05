package day3

import (
	"adventofcode/utils"
	"math"
	"strconv"
)

type BinaryDiagnostic struct{}

func (BinaryDiagnostic) Input() string {
	return "2021/03_binary_diagnostic/input.txt"
}

func gamma_epsilon(arr []int) (gamma int, epsilon int) {
	for i, b := range arr {
		digit := len(arr) - i - 1
		if b > 0 {
			gamma += int(math.Pow(2, float64(digit)))
		} else {
			epsilon += int(math.Pow(2, float64(digit)))
		}

	}
	return gamma, epsilon
}

func determineWeight(rl utils.ReadLine) []int {
	line, eof := rl()
	arr := make([]int, len(line))
	for !eof {
		for i, b := range line {
			if b == '0' {
				arr[i]--
			} else {
				arr[i]++
			}
		}
		line, eof = rl()
	}
	return arr
}

func (BinaryDiagnostic) SolveQ1(rl utils.ReadLine) string {
	arr := determineWeight(rl)

	gamma, epsilon := gamma_epsilon(arr)

	return strconv.Itoa(gamma * epsilon)
}

func generateBitArr(rl utils.ReadLine) [][]bool {
	line, eof := rl()
	var arrs [][]bool
	for !eof {
		arr := make([]bool, len(line))
		for i, b := range line {
			if b == '0' {
				arr[i] = false
			} else {
				arr[i] = true
			}
		}
		arrs = append(arrs, arr)
		line, eof = rl()
	}
	return arrs
}

func o2(digit int, arrs [][]bool) []bool {
	if len(arrs) == 1 {
		return arrs[0]
	}

	var zeroes [][]bool
	var ones [][]bool
	for _, arr := range arrs {
		if arr[digit] {
			ones = append(ones, arr)
		} else {
			zeroes = append(zeroes, arr)
		}
	}
	if len(ones) >= len(zeroes) {
		return o2(digit+1, ones)
	}
	return o2(digit+1, zeroes)
}

func co2(digit int, arrs [][]bool) []bool {
	if len(arrs) == 1 {
		return arrs[0]
	}

	var zeroes [][]bool
	var ones [][]bool
	for _, arr := range arrs {
		if arr[digit] {
			ones = append(ones, arr)
		} else {
			zeroes = append(zeroes, arr)
		}
	}
	if len(zeroes) <= len(ones) {
		return co2(digit+1, zeroes)
	}
	return co2(digit+1, ones)
}

func toInt(arr []bool) int {
	acc := 0
	for i, b := range arr {
		digit := len(arr) - i - 1
		if b {
			acc += int(math.Pow(2, float64(digit)))
		}
	}
	return acc
}

func (BinaryDiagnostic) SolveQ2(rl utils.ReadLine) string {
	arrs := generateBitArr(rl)
	o2 := toInt(o2(0, arrs))
	co2 := toInt(co2(0, arrs))

	return strconv.Itoa(o2 * co2)
}
