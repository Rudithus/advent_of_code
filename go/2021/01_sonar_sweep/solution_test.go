package day1

import (
	"adventofcode/utils/files"
	"testing"
)

var problem = Sonarsweep{}

func TestQ1(t *testing.T) {
	lineReader := files.LineReader("test.txt")
	expectedAnswer := "7"

	answer := problem.SolveQ1(lineReader)
	if expectedAnswer != answer {
		t.Error("expected:" + expectedAnswer + " got:" + answer)
	}
}

func TestQ2(t *testing.T) {
	lineReader := files.LineReader("test.txt")
	expectedAnswer := "5"
	answer := problem.SolveQ2(lineReader)

	if expectedAnswer != answer {
		t.Error("expected:" + expectedAnswer + " got:" + answer)

	}
}
