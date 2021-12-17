package day11

import (
	"adventofcode/tests"
	"testing"
)

func TestSolveQ1(t *testing.T) {
	tests.TestQ1(t, DumboOctopus{}, 1656)
}

func TestSolveQ2(t *testing.T) {
	tests.TestQ2(t, DumboOctopus{}, 195)
}
