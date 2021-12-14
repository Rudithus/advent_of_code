package day9

import (
	"adventofcode/tests"
	"testing"
)

func TestSolveQ1(t *testing.T) {
	tests.TestQ1(t, SmokeBasin{}, 15)
}

func TestSolveQ2(t *testing.T) {
	tests.TestQ2(t, SmokeBasin{}, 1134)
}
