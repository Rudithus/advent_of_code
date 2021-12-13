package main

import (
	twentytwentytwo "adventofcode/2021"
	"adventofcode/utils"
	"fmt"
	"os"
)

var problems = twentytwentytwo.Problems()

func main() {
	problem := problems[os.Args[1]]
	question := os.Args[2]
	input, err := os.ReadFile(problem.Path() + "/input.txt")

	utils.Check(err)

	if question == "q2" {
		fmt.Println(problem.SolveQ2(input))
		return
	}

	fmt.Println(problem.SolveQ1(input))
}
