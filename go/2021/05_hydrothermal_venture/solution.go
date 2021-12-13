package day5

import (
	"adventofcode/utils"
	"bufio"
	"strconv"
	"strings"
)

type HydrothermalVent struct{}

func (HydrothermalVent) Path() string { return "2021/05_hydrothermal_venture" }

type Coordinate struct {
	X int
	Y int
}

type Vent struct {
	Start Coordinate
	End   Coordinate
}

type Grid struct {
	arr         [1000][1000]int
	dangerCount int
}

func (origin *Coordinate) traverser(dest Coordinate) func() (Coordinate, bool) {
	pointer := Coordinate{origin.X, origin.Y}
	arrived := false

	return func() (Coordinate, bool) {
		if !arrived {
			pointer.X += utils.CompareInt(dest.X, pointer.X)
			pointer.Y += utils.CompareInt(dest.Y, pointer.Y)

			arrived = pointer.X == dest.X && pointer.Y == dest.Y
		}
		return pointer, arrived
	}
}

func coordinate(c string) Coordinate {
	arr := strings.Split(c, ",")
	x, err := strconv.Atoi(arr[0])
	utils.Check(err)
	y, err := strconv.Atoi(arr[1])
	utils.Check(err)

	return Coordinate{X: x, Y: y}
}

func vent(line string) Vent {
	segments := strings.Split(line, " -> ")
	begining := coordinate(segments[0])
	end := coordinate(segments[1])
	return Vent{begining, end}
}

func (g *Grid) markCoordinate(c Coordinate) {
	g.arr[c.Y][c.X]++
	if g.arr[c.Y][c.X] == 2 {
		g.dangerCount++
	}
}
func (HydrothermalVent) SolveQ1(input []byte) int {
	var grid Grid
	read := utils.Reader(input, bufio.ScanLines)
	for line, eof := read(); !eof; line, eof = read() {
		vent := vent(line)
		if vent.Start.X == vent.End.X || vent.Start.Y == vent.End.Y {
			grid.markCoordinate(vent.Start)
			traverse := vent.Start.traverser(vent.End)
			for {
				coord, arrived := traverse()
				grid.markCoordinate(coord)
				if arrived {
					break
				}
			}
		}
	}
	return grid.dangerCount
}

func (HydrothermalVent) SolveQ2(input []byte) int {
	var grid Grid
	read := utils.Reader(input, bufio.ScanLines)

	for line, eof := read(); !eof; line, eof = read() {
		vent := vent(line)
		grid.markCoordinate(vent.Start)
		traverse := vent.Start.traverser(vent.End)
		for {
			coord, arrived := traverse()
			grid.markCoordinate(coord)
			if arrived {
				break
			}
		}
	}
	return grid.dangerCount
}
