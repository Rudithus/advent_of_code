package day4

import (
	"adventofcode/utils/files"
	"testing"
)

var problem = GiantSquid{}

func TestQ1(t *testing.T) {
	lineReader := files.LineReader("test.txt")
	expectedAnswer := "4512"
	answer := problem.SolveQ1(lineReader)

	if expectedAnswer != answer {
		t.Error("expected:" + expectedAnswer + " got:" + answer)
	}
}

func TestQ2(t *testing.T) {
	lineReader := files.LineReader("test.txt")
	expectedAnswer := "1924"
	answer := problem.SolveQ2(lineReader)

	if expectedAnswer != answer {
		t.Error("expected:" + expectedAnswer + " got:" + answer)
	}
}
