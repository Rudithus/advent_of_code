package day2

import (
	"adventofcode/tests"
	"testing"
)

func TestSolveQ1(t *testing.T) {
	tests.TestQ1(t, Dive{}, 150)
}

func TestSolveQ2(t *testing.T) {
	tests.TestQ2(t, Dive{}, 900)
}
