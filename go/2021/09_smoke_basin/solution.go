package day9

import (
	"adventofcode/utils"
	"bufio"
	"strconv"
)

type SmokeBasin struct{}

func (SmokeBasin) Path() string {
	return "2021/09_smoke_basin"
}

type location struct {
	height int
	x, y   int
	grid   *[][]location
}

func (loc *location) adjLocations() (adjs []location) {
	grid := *loc.grid
	x, y := loc.x, loc.y
	if x != 0 {
		adjs = append(adjs, grid[y][x-1])
	}
	if y != 0 {
		adjs = append(adjs, grid[y-1][x])
	}
	if y != len(grid)-1 {
		adjs = append(adjs, grid[y+1][x])
	}
	if x != len(grid[0])-1 {
		adjs = append(adjs, grid[y][x+1])
	}
	return
}
func (loc *location) isLowPoint() bool {
	result := true
	for _, l := range loc.adjLocations() {
		result = result && loc.height < l.height
	}
	return result
}
func (loc *location) riskLevel() int {
	return loc.height + 1
}

func (loc *location) basin(visited *map[location]utils.Empty) *map[location]utils.Empty {
	adjs := loc.adjLocations()
	tt := *visited
	for _, l := range adjs {
		if _, visit := tt[l]; !visit && l.height != 9 {
			tt[l] = utils.Empty{}
			l.basin(visited)
		}
	}
	return visited
}

func prepareGrid(read func() (string, bool)) (grid [][]location) {
	var locations []location
	x, y := 0, 0
	for ch, eoi := read(); !eoi; ch, eoi = read() {
		if ch == "\n" {
			grid = append(grid, locations)
			locations = []location{}
			y++
			x = 0
			continue
		}

		h, _ := strconv.Atoi(ch)
		locations = append(locations, location{h, x, y, &grid})
		x++
	}
	grid = append(grid, locations)
	return
}

func (SmokeBasin) SolveQ1(input []byte) int {
	riskLevel := 0
	grid := prepareGrid(utils.Reader(input, bufio.ScanRunes))
	for _, lines := range grid {
		for _, location := range lines {
			if location.isLowPoint() {
				riskLevel += location.riskLevel()
			}
		}
	}
	return riskLevel
}

func (SmokeBasin) SolveQ2(input []byte) int {
	orderedArray := utils.OrderedIntArray{0}
	grid := prepareGrid(utils.Reader(input, bufio.ScanRunes))
	for _, lines := range grid {
		for _, loc := range lines {
			if loc.isLowPoint() {
				visited := make(map[location]utils.Empty)
				basin := loc.basin(&visited)
				orderedArray.Append(len(*basin))
			}
		}
	}
	length := len(orderedArray)
	return orderedArray[length-1] * orderedArray[length-2] * orderedArray[length-3]
}
