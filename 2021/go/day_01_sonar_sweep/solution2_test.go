package day1

import (
	"os"
	"testing"
)

func TestSolve2(t *testing.T) {
	file, err := os.ReadFile("input.txt")
	expectedAnswer := 5

	if err != nil {
		return
	}

	if expectedAnswer != Solve2(string(file)) {
		t.Fail()
	}
}
