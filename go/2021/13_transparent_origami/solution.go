package day13

import (
	"adventofcode/utils"
	"bufio"
	"strconv"
	"strings"
)

type TransparentOrigami struct{}

func (TransparentOrigami) Path() string {
	return "2021/13_transparent_origami"
}

type grid [][]bool

func parseInput(input []byte) (g grid, folds []fold) {
	read := utils.Reader(input, bufio.ScanLines)
	maxX := 0
	for line, _ := read(); line != ""; line, _ = read() {
		arr := strings.Split(line, ",")
		x, _ := strconv.Atoi(arr[0])
		y, _ := strconv.Atoi(arr[1])

		if len(g) <= y {
			g = append(g, make([][]bool, 1+(y-len(g)))...)
		}
		if len(g[y]) <= x {
			if x > maxX {
				maxX = x
			}
			g[y] = append(g[y], make([]bool, 1+(x-len(g[y])))...)
		}
		g[y][x] = true
	}

	for y, arr := range g {
		if len(arr) != maxX+1 {
			g[y] = append(arr, make([]bool, 1+(maxX-len(g[y])))...)
		}
	}

	for line, eoi := read(); !eoi; line, eoi = read() {
		axis := line[11:12]
		mp, _ := strconv.Atoi(line[13:])
		f := fold{midPoint: mp}
		if axis == "y" {
			f.vertical = true
			folds = append(folds, f)
		} else {
			f.vertical = false
			folds = append(folds, f)
		}
	}

	return
}

func (g grid) countDots() int {
	count := 0
	for _, arr := range g {
		for _, d := range arr {
			if d {
				count++
			}
		}
	}
	return count
}

type fold struct {
	vertical bool
	midPoint int
}

func (g grid) verticalFold(midPoint int) (result grid) {
	one, other := g[:midPoint], g[midPoint+1:]
	if len(one) > len(other) {
		diff := len(one) - len(other)
		patch := make(grid, diff)
		for i := range patch {
			patch[i] = make([]bool, len(one[0]))
		}
		other = append(other, patch...)
	} else if len(other) > len(one) {
		diff := len(other) - len(one)
		patch := make(grid, diff)
		for i := range patch {
			patch[i] = make([]bool, len(one[0]))
		}

		one = append(patch, one...)
	}
	result = make(grid, len(one))
	for y := 0; y < len(one); y++ {
		otherY := len(one) - y - 1
		for x := 0; x < len(one[y]); x++ {
			one[y][x] = one[y][x] || other[otherY][x]
		}
		result[y] = one[y]
	}
	return
}

func (g grid) horizontalFold(midPoint int) (result grid) {
	result = make(grid, len(g))
	for y, arr := range g {
		one, other := arr[:midPoint], arr[midPoint+1:]

		if len(one) > len(other) {
			diff := len(one) - len(other)
			other = append(other, make([]bool, diff)...)
		} else if len(other) > len(one) {
			diff := len(other) - len(one)
			one = append(make([]bool, diff), one...)
		}

		for x := 0; x < len(one); x++ {
			otherX := len(one) - x - 1
			one[x] = one[x] || other[otherX]
		}
		result[y] = one
	}
	return
}

func (TransparentOrigami) SolveQ1(input []byte) int {
	grid, folds := parseInput(input)
	f := folds[0]
	if f.vertical {
		grid = grid.verticalFold(f.midPoint)
	} else {
		grid = grid.horizontalFold(f.midPoint)
	}

	return grid.countDots()
}

func (TransparentOrigami) SolveQ2(input []byte) int {
	grid, folds := parseInput(input)
	for _, f := range folds {
		if f.vertical {
			grid = grid.verticalFold(f.midPoint)
		} else {
			grid = grid.horizontalFold(f.midPoint)
		}
	}
	utils.PrintGrid(grid)
	return grid.countDots()
}
