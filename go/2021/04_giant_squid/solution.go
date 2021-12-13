package day4

import (
	"adventofcode/utils"
	"bufio"
	"strconv"
	"strings"
)

type GiantSquid struct{}

func (GiantSquid) Path() string {
	return "2021/04_giant_squid"
}

func prepareNumbers() (numbers [100]*BingoNumber) {
	for i := 0; i < 100; i++ {
		numbers[i] = &BingoNumber{Value: i, Drawn: false}
	}
	return
}

func getDraw(read func() (string, bool)) []string {
	line, _ := read()
	return strings.Split(line, ",")
}

func createBoard(read func() (string, bool), bns [100]*BingoNumber) bool {
	bb := &BingoBoard{Complete: false}
	var verticalSeries [5]*BingoSeries
	for i := 0; i < 5; i++ {
		verticalSeries[i] = &BingoSeries{BingoBoard: bb}
	}

	lineCount := 0
	for {
		line, eof := read()
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

func prepareBingo(read func() (string, bool)) (numbers [100]*BingoNumber, draw []string, boardCount int) {
	numbers = prepareNumbers()
	draw = getDraw(read)
	read()
	boardCount = 1
	for !createBoard(read, numbers) {
		boardCount++
	}

	return
}

func (GiantSquid) SolveQ1(input []byte) int {
	read := utils.Reader(input, bufio.ScanLines)
	numbers, draw, _ := prepareBingo(read)

	for _, char := range draw {
		value, _ := strconv.Atoi(char)
		number := numbers[value]
		winners := number.Mark()
		if len(winners) > 0 {
			return winners[0].Score
		}
	}
	return 0
}

func (GiantSquid) SolveQ2(input []byte) int {
	read := utils.Reader(input, bufio.ScanLines)
	numbers, draw, _ := prepareBingo(read)
	var lastWinner *BingoBoard

	for _, char := range draw {
		value, _ := strconv.Atoi(char)
		number := numbers[value]
		winners := number.Mark()
		if len(winners) > 0 {
			lastWinner = winners[len(winners)-1]
		}
	}

	return lastWinner.Score
}
