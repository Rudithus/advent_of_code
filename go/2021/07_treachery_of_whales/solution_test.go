package day7

import (
	"adventofcode/utils/files"
	"testing"
)

var problem = TreacheryOfWhales{}

func TestQ1(t *testing.T) {
	lineReader := files.CharReader("test.txt")
	expectedAnswer := "37"

	answer := problem.SolveQ1(lineReader)
	if expectedAnswer != answer {
		t.Error("expected:" + expectedAnswer + " got:" + answer)
	}
}

func TestQ2(t *testing.T) {
	lineReader := files.CharReader("test.txt")
	expectedAnswer := "168"
	answer := problem.SolveQ2(lineReader)

	if expectedAnswer != answer {
		t.Error("expected:" + expectedAnswer + " got:" + answer)

	}
}
