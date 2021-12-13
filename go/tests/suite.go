package tests

import (
	"adventofcode/utils"
	"os"
	"testing"
)

func input() []byte {
	input, err := os.ReadFile("test.txt")
	utils.Check(err)
	return input
}
func TestQ1(t *testing.T, problem utils.Problem, expected int) {
	actual := problem.SolveQ1(input())
	if actual != expected {
		t.Fatalf("expected:%d, got: %d", expected, actual)
	}
}

func TestQ2(t *testing.T, problem utils.Problem, expected int) {
	actual := problem.SolveQ2(input())
	if actual != expected {
		t.Fatalf("expected:%d, got: %d", expected, actual)
	}
}
