package day13

import (
	"adventofcode/tests"
	"testing"
)

func TestSolveQ1(t *testing.T) {
	tests.TestQ1(t, TransparentOrigami{}, 17)
}

func TestSolveQ2(t *testing.T) {
	tests.TestQ2(t, TransparentOrigami{}, 16)
}
