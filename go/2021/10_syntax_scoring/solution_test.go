package day10

import (
	"adventofcode/tests"
	"testing"
)

func TestSolveQ1(t *testing.T) {
	tests.TestQ1(t, SyntaxScoring{}, 26397)
}

func TestSolveQ2(t *testing.T) {
	tests.TestQ2(t, SyntaxScoring{}, 288957)
}
