package day6

import (
	"adventofcode/tests"
	"testing"
)

func TestSolveQ1(t *testing.T) {
	tests.TestQ1(t, Lanternfish{}, 5934)
}

func TestSolveQ2(t *testing.T) {
	tests.TestQ2(t, Lanternfish{}, 26984457539)
}
