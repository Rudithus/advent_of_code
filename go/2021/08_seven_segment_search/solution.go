package day8

import (
	"adventofcode/utils"
	"bufio"
	"strconv"
	"strings"
)

type SevenSegment struct{}

func (SevenSegment) Path() string {
	return "2021/08_seven_segment_search"
}

func (SevenSegment) SolveQ1(input []byte) int {
	counter := 0
	read := utils.Reader(input, bufio.ScanLines)
	for line, eof := read(); !eof; line, eof = read() {
		entry := strings.Fields(line)
		output := entry[11:]
		for _, v := range output {
			switch len(v) {
			case 2:
				fallthrough
			case 4:
				fallthrough
			case 3:
				fallthrough
			case 7:
				counter++
			}
		}
	}

	return counter
}

func easyFour(patterns []string) (one, four, seven, eight int, rest []int) {
	for _, pattern := range patterns {
		byteArr := []byte(pattern)
		digit := utils.MultiplyBytes(byteArr)
		switch len(pattern) {
		case 2:
			one = digit
		case 4:
			four = digit
		case 3:
			seven = digit
		case 7:
			eight = digit
		default:
			rest = append(rest, digit)
		}
	}
	return
}

func generateCypher(patterns []string) map[int]int {
	one, four, seven, eight, rest := easyFour(patterns)
	cypher := map[int]int{}
	cypher[one] = 1
	cypher[four] = 4
	cypher[seven] = 7
	cypher[eight] = 8

	top := seven / one
	middleXtopLeft := four / one
	bottomXbottomLeft := eight / (four * top)
	var topRight, middle, bottom int
	for _, digit := range rest {
		leftOver := eight / digit
		if leftOver < 120 {
			if one%leftOver == 0 {
				cypher[digit] = 6
				topRight = leftOver
				continue
			}
			if middleXtopLeft%leftOver == 0 {
				cypher[digit] = 0
				middle = leftOver
				continue
			}
			if bottomXbottomLeft%leftOver == 0 {
				cypher[digit] = 9
				bottom = digit / (four * top)
				continue
			}
		}
	}

	bottomRight := one / topRight
	topLeft := middleXtopLeft / middle
	two := topRight * middle * bottomXbottomLeft * top
	three := bottom * middle * top * one
	five := bottom * middle * top * bottomRight * topLeft
	cypher[two] = 2
	cypher[three] = 3
	cypher[five] = 5
	return cypher
}
func processEntry(entry []string) int {
	patterns := entry[:10]
	cypher := generateCypher(patterns)
	output := entry[11:]

	result := ""
	for _, signal := range output {
		sum := utils.MultiplyBytes([]byte(signal))
		c := cypher[sum]
		result += strconv.Itoa(c)
	}

	val, _ := strconv.Atoi(result)
	return val
}
func (SevenSegment) SolveQ2(input []byte) int {
	read := utils.Reader(input, bufio.ScanLines)
	total := 0
	for line, eof := read(); !eof; line, eof = read() {
		r := processEntry(strings.Fields(line))
		total += r
	}
	return total
}
