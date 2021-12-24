package day12

import (
	"adventofcode/tests"
	"testing"
)

func TestSolveQ1(t *testing.T) {
	tests.TestQ1(t, PassagePathing{}, 226)
}

func TestSolveQ2(t *testing.T) {
	tests.TestQ2(t, PassagePathing{}, 3509)
}
