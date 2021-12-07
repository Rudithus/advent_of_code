package day4

import (
	"adventofcode/utils"
	"strconv"
	"strings"
)

type GiantSquid struct{}

func (GiantSquid) Input() string {
	return "2021/04_giant_squid/input.txt"
}

func prepareNumbers() (numbers [100]*BingoNumber) {
	for i := 0; i < 100; i++ {
		numbers[i] = &BingoNumber{Value: i, Drawn: false}
	}
	return
}

func getDraw(rl utils.ReadLine) []string {
	line, _ := rl()
	return strings.Split(line, ",")
}

func createBoard(rl utils.ReadLine, bns [100]*BingoNumber) bool {
	bb := &BingoBoard{Complete: false}
	var verticalSeries [5]*BingoSeries
	for i := 0; i < 5; i++ {
		verticalSeries[i] = &BingoSeries{BingoBoard: bb}
	}

	lineCount := 0
	for {
		line, eof := rl()
		if eof || line == "" {
			return eof
		}
		bs := BingoSeries{BingoBoard: bb}
		fields := strings.Fields(line)
		for i, char := range fields {
			value, _ := strconv.Atoi(char)
			bn := bns[value]

			bb.BingoNumbers[i+(lineCount*5)] = bn
			bs.BingoNumbers[i] = bn
			verticalSeries[i].BingoNumbers[lineCount] = bn

			bn.BingoSeries = append(bn.BingoSeries, &bs)
			bn.BingoSeries = append(bn.BingoSeries, verticalSeries[i])
		}
		lineCount++
	}
}

func prepareBingo(rl utils.ReadLine) (numbers [100]*BingoNumber, draw []string, boardCount int) {
	numbers = prepareNumbers()
	draw = getDraw(rl)
	rl()
	boardCount = 1
	for !createBoard(rl, numbers) {
		boardCount++
	}

	return
}

func (GiantSquid) SolveQ1(rl utils.ReadLine) string {
	numbers, draw, _ := prepareBingo(rl)

	for _, char := range draw {
		value, _ := strconv.Atoi(char)
		number := numbers[value]
		winners := number.Mark()
		if len(winners) > 0 {
			return strconv.Itoa(winners[0].Score)
		}
	}
	return ""
}

func (GiantSquid) SolveQ2(rl utils.ReadLine) string {
	numbers, draw, _ := prepareBingo(rl)
	var lastWinner *BingoBoard
	for _, char := range draw {
		value, _ := strconv.Atoi(char)
		number := numbers[value]
		winners := number.Mark()
		if len(winners) > 0 {
			lastWinner = winners[len(winners)-1]
		}
	}

	return strconv.Itoa(lastWinner.Score)
}
