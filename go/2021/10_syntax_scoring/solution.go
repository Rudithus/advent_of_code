package day10

import (
	"adventofcode/utils"
	"bufio"
)

type SyntaxScoring struct{}

var brackets = map[string]string{")": "(", "}": "{", "]": "[", ">": "<"}

func (SyntaxScoring) Path() string {
	return "2021/10_syntax_scoring"
}

func (SyntaxScoring) SolveQ1(input []byte) int {
	totalScore := 0
	read := utils.Reader(input, bufio.ScanBytes)
	scoreTable := map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}
	stack := utils.NewStringStack()
	for ch, eoi := read(); !eoi; ch, eoi = read() {
		if ch == "\n" {
			stack = utils.NewStringStack()
			continue
		}

		if match, isClosing := brackets[ch]; isClosing {
			if pop := stack.Pop(); pop != match {
				totalScore += scoreTable[ch]
				continue
			}
			continue
		}
		stack.Push(ch)
	}
	return totalScore
}
func readUntilNewLine(read func() (string, bool)) {
	for ch, _ := read(); ch != "\n"; ch, _ = read() {
	}
}
func scoreCompletion(stack utils.StringStack) int {
	totalScore := 0
	scoreTable := map[string]int{")": 1, "]": 2, "}": 3, ">": 4}
	brackets := map[string]string{"(": ")", "{": "}", "[": "]", "<": ">"}
	for i := 0; !stack.Empty(); i++ {
		totalScore = totalScore * 5
		pop := stack.Pop()
		match := brackets[pop]
		totalScore += scoreTable[match]
	}
	return totalScore
}
func (SyntaxScoring) SolveQ2(input []byte) int {
	scores := utils.OrderedIntArray{0}
	read := utils.Reader(input, bufio.ScanBytes)
	stack := utils.NewStringStack()
	for ch, eoi := read(); !eoi; ch, eoi = read() {
		if ch == "\n" {
			scores.Append(scoreCompletion(stack))
			stack = utils.NewStringStack()
			continue
		}

		if match, isClosing := brackets[ch]; isClosing {
			if pop := stack.Pop(); pop != match {
				//corrupted
				readUntilNewLine(read)
				stack = utils.NewStringStack()
				continue
			}
			continue
		}
		stack.Push(ch)
	}
	scores.Append(scoreCompletion(stack))
	midIndex := utils.Middle(scores)
	return scores[midIndex]
}
