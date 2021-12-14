package tests

import (
	twentytwentytwo "adventofcode/2021"
	"adventofcode/utils"
	"os"
	"reflect"
	"testing"
)

func BenchmarkSolutions(b *testing.B) {
	for _, p := range twentytwentytwo.Problems() {
		testName := reflect.TypeOf(p).Name()
		input, err := os.ReadFile("../" + p.Path() + "/input.txt")
		utils.Check(err)
		b.ResetTimer()
		b.Run(testName+"Q1", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				p.SolveQ1(input)
			}
		})
		b.Run(testName+"Q2", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				p.SolveQ2(input)
			}
		})
	}
}
