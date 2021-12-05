package day1

import (
	"os"
	"testing"
)

func TestSolve(t *testing.T) {
	file, err := os.ReadFile("input.txt")
	expectedAnswer := 1016964

	if err != nil {
		return
	}

	if expectedAnswer != solve(string(file)) {
		t.Fail()
	}
}
