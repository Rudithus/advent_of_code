package main

import (
	day1 "adventofcode/2021/01_sonar_sweep"
	day2 "adventofcode/2021/02_dive"
	day3 "adventofcode/2021/03_binary_diagnostic"
	day4 "adventofcode/2021/04_giant_squid"
	day5 "adventofcode/2021/05_hydrothermal_venture"
	day6 "adventofcode/2021/06_lanternfish"
	day7 "adventofcode/2021/07_treachery_of_whales"
	"adventofcode/utils"
	"adventofcode/utils/files"
	"fmt"
	"os"
)

var problems = map[string]utils.Problem{
	"sonarsweep":  day1.Sonarsweep{},
	"dive":        day2.Dive{},
	"binary":      day3.BinaryDiagnostic{},
	"bingo":       day4.GiantSquid{},
	"hydro":       day5.HydrothermalVent{},
	"lanternfish": day6.Lanternfish{},
	"whales":      day7.TreacheryOfWhales{},
}

func main() {
	problem := problems[os.Args[1]]
	question := os.Args[2]

	linereader := files.CharReader(problem.Input())
	if question == "q2" {
		fmt.Println(problem.SolveQ2(linereader))
		return
	}

	fmt.Println(problem.SolveQ1(linereader))
}
