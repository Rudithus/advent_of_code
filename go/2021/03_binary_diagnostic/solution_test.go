package day3

import (
	"adventofcode/utils/files"
	"testing"
)

var problem = BinaryDiagnostic{}

func TestQ1(t *testing.T) {
	lineReader := files.LineReader("test.txt")
	expectedAnswer := "198"
	answer := problem.SolveQ1(lineReader)

	if expectedAnswer != answer {
		t.Error("expected:" + expectedAnswer + " got:" + answer)
	}
}

func TestQ2(t *testing.T) {
	lineReader := files.LineReader("test.txt")
	expectedAnswer := "230"
	answer := problem.SolveQ2(lineReader)

	if expectedAnswer != answer {
		t.Error("expected:" + expectedAnswer + " got:" + answer)
	}
}
