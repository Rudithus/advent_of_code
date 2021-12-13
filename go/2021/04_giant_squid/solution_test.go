package day4

import (
	"adventofcode/tests"
	"testing"
)

func TestSolveQ1(t *testing.T) {
	tests.TestQ1(t, GiantSquid{}, 4512)
}

func TestSolveQ2(t *testing.T) {
	tests.TestQ2(t, GiantSquid{}, 1924)
}
