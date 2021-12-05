package day2

import (
	"adventofcode/utils/files"
	"testing"
)

var problem = Dive{}

func TestQ1(t *testing.T) {
	lineReader := files.LineReader("test.txt")
	expectedAnswer := "150"
	answer := problem.SolveQ1(lineReader)
	if expectedAnswer != answer {
		t.Error("expected:" + expectedAnswer + " got:" + answer)
	}
}

func TestQ2(t *testing.T) {
	lineReader := files.LineReader("test.txt")
	expectedAnswer := "900"
	if expectedAnswer != problem.SolveQ2(lineReader) {
		t.Fail()
	}
}
