package day8

import (
	"adventofcode/tests"
	"testing"
)

func TestSolveQ1(t *testing.T) {
	tests.TestQ1(t, SevenSegment{}, 26)
}

func TestSolveQ2(t *testing.T) {
	tests.TestQ2(t, SevenSegment{}, 61229)
}
