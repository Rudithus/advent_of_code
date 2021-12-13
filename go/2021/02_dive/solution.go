package day2

import (
	"adventofcode/2021/02_dive/q1"
	"adventofcode/2021/02_dive/q2"
	"adventofcode/utils"
	"bufio"
	"strconv"
	"strings"
)

type Dive struct{}

func (Dive) Path() string {
	return "2021/02_dive"
}

type Submarine interface {
	Forward(value int)
	Down(value int)
	Up(value int)
}

func process(sub Submarine, input []byte) {
	read := utils.Reader(input, bufio.ScanLines)
	for line, eof := read(); !eof; line, eof = read() {
		fields := strings.Fields(line)
		direction := fields[0]
		value, _ := strconv.Atoi(fields[1])
		switch direction {
		case "forward":
			sub.Forward(value)
		case "up":
			sub.Up(value)
		case "down":
			sub.Down(value)
		}
	}
}
func (Dive) SolveQ1(input []byte) int {
	sub := q1.Submarine{Pos: q1.Position{X: 0, Y: 0}}
	process(&sub, input)
	return sub.Pos.X * sub.Pos.Y
}

func (Dive) SolveQ2(input []byte) int {
	sub := q2.Submarine{Pos: q2.Position{X: 0, Y: 0}, Aim: 0}
	process(&sub, input)
	return sub.Pos.X * sub.Pos.Y
}
