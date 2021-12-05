package day2

import (
	"adventofcode/2021/02_dive/q1"
	"adventofcode/2021/02_dive/q2"
	"adventofcode/utils"
	"strconv"
	"strings"
)

type Dive struct{}

func (Dive) Input() string {
	return "2021/02_dive/input.txt"
}

type Submarine interface {
	Forward(value int)
	Down(value int)
	Up(value int)
}

func process(sub Submarine, rl utils.ReadLine) {
	for {
		line, eof := rl()
		if eof {
			break
		}

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
func (Dive) SolveQ1(rl utils.ReadLine) string {
	sub := q1.Submarine{Pos: q1.Position{X: 0, Y: 0}}
	process(&sub, rl)
	return strconv.Itoa(sub.Pos.X * sub.Pos.Y)
}

func (Dive) SolveQ2(rl utils.ReadLine) string {
	sub := q2.Submarine{Pos: q2.Position{X: 0, Y: 0}, Aim: 0}
	process(&sub, rl)
	return strconv.Itoa(sub.Pos.X * sub.Pos.Y)
}
