package main

import (
	day1 "adventofcode/2021/01_sonar_sweep"
	day2 "adventofcode/2021/02_dive"
	"adventofcode/utils"
	"adventofcode/utils/files"
	"fmt"
	"os"
)

var problems = map[string]utils.Problem{
	"sonarsweep": day1.Sonarsweep{},
	"dive":       day2.Dive{},
}

func main() {
	problem := problems[os.Args[1]]
	question := os.Args[2]

	linereader := files.LineReader(problem.Input())
	if question == "q2" {
		fmt.Println(problem.SolveQ2(linereader))
		return
	}

	fmt.Println(problem.SolveQ1(linereader))
}
