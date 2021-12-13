package day3

import (
	"adventofcode/tests"
	"testing"
)

func TestSolveQ1(t *testing.T) {
	tests.TestQ1(t, BinaryDiagnostic{}, 198)
}

func TestSolveQ2(t *testing.T) {
	tests.TestQ2(t, BinaryDiagnostic{}, 230)
}
