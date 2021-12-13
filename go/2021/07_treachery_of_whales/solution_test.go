package day7

import (
	"adventofcode/tests"
	"testing"
)

func TestSolveQ1(t *testing.T) {
	tests.TestQ1(t, TreacheryOfWhales{}, 37)
}

func TestSolveQ2(t *testing.T) {
	tests.TestQ2(t, TreacheryOfWhales{}, 168)
}
