package day1

import (
	"adventofcode/tests"
	"testing"
)

func TestSolveQ1(t *testing.T) {
	tests.TestQ1(t, Sonarsweep{}, 7)
}

func TestSolveQ2(t *testing.T) {
	tests.TestQ2(t, Sonarsweep{}, 5)
}
