package day5

import (
	"adventofcode/tests"
	"testing"
)

func TestSolveQ1(t *testing.T) {
	tests.TestQ1(t, HydrothermalVent{}, 5)
}

func TestSolveQ2(t *testing.T) {
	tests.TestQ2(t, HydrothermalVent{}, 12)
}
