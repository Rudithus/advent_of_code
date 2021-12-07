package day4

type BingoNumber struct {
	Value       int
	Drawn       bool
	BingoSeries []*BingoSeries
}

type BingoSeries struct {
	BingoNumbers [5]*BingoNumber
	BingoBoard   *BingoBoard
}

type BingoBoard struct {
	BingoNumbers [25]*BingoNumber
	Complete     bool
	Score        int
}

func (bn *BingoNumber) Mark() (winners []*BingoBoard) {
	bn.Drawn = true
	for _, bs := range bn.BingoSeries {
		if bs.BingoBoard.Complete {
			continue
		}
		if bs.Check() {
			bs.BingoBoard.Score = calculateScore(bs.BingoBoard) * bn.Value
			winners = append(winners, bs.BingoBoard)
		}
	}
	return
}
func (bs *BingoSeries) Check() bool {
	for _, n := range bs.BingoNumbers {
		if !n.Drawn {
			return false
		}
	}
	bs.BingoBoard.Complete = true
	return true
}

func calculateScore(bb *BingoBoard) int {
	acc := 0
	for _, n := range bb.BingoNumbers {
		if n.Drawn {
			continue
		}
		acc += n.Value
	}
	return acc
}
