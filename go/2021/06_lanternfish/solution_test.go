package day6

import (
	"adventofcode/utils/files"
	"testing"
)

var problem = Lanternfish{}

func TestQ1(t *testing.T) {
	lineReader := files.LineReader("test.txt")
	expectedAnswer := "5934"

	answer := problem.SolveQ1(lineReader)
	if expectedAnswer != answer {
		t.Error("expected:" + expectedAnswer + " got:" + answer)
	}
}

func TestQ2(t *testing.T) {
	lineReader := files.LineReader("test.txt")
	expectedAnswer := "26984457539"
	answer := problem.SolveQ2(lineReader)

	if expectedAnswer != answer {
		t.Error("expected:" + expectedAnswer + " got:" + answer)

	}
}
