package twentytwentytwo

import (
	day1 "adventofcode/2021/01_sonar_sweep"
	day2 "adventofcode/2021/02_dive"
	day3 "adventofcode/2021/03_binary_diagnostic"
	day4 "adventofcode/2021/04_giant_squid"
	day5 "adventofcode/2021/05_hydrothermal_venture"
	day6 "adventofcode/2021/06_lanternfish"
	day7 "adventofcode/2021/07_treachery_of_whales"
	day8 "adventofcode/2021/08_seven_segment_search"
	day9 "adventofcode/2021/09_smoke_basin"
	day10 "adventofcode/2021/10_syntax_scoring"
	day11 "adventofcode/2021/11_dumbo_octopus"
	day12 "adventofcode/2021/12_passage_pathing"
	day13 "adventofcode/2021/13_transparent_origami"
	"adventofcode/utils"
)

var problems = map[string]utils.Problem{
	"sonarsweep":  day1.Sonarsweep{},
	"dive":        day2.Dive{},
	"binary":      day3.BinaryDiagnostic{},
	"bingo":       day4.GiantSquid{},
	"hydro":       day5.HydrothermalVent{},
	"lanternfish": day6.Lanternfish{},
	"whales":      day7.TreacheryOfWhales{},
	"seven":       day8.SevenSegment{},
	"smoke":       day9.SmokeBasin{},
	"syntax":      day10.SyntaxScoring{},
	"dumbo":       day11.DumboOctopus{},
	"passage":     day12.PassagePathing{},
	"origami":     day13.TransparentOrigami{},
}

func Problems() map[string]utils.Problem {
	return problems
}
