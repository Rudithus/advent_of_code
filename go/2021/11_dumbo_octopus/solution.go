package day11

import (
	"adventofcode/utils"
	"bufio"
	"strconv"
)

type DumboOctopus struct{}

func (DumboOctopus) Path() string {
	return "2021/11_dumbo_octopus"
}

type octo struct {
	energyLevel int
	exhausted   bool
	neighbours  []*octo
}

func (o *octo) flash() (totalFlashers []*octo) {
	if o.exhausted {
		return
	}
	o.exhausted = true
	totalFlashers = append(totalFlashers, o)
	for _, adj := range o.neighbours {
		adj.increaseLevel()
		if adj.energyLevel > 9 {
			totalFlashers = append(totalFlashers, adj.flash()...)
		}
	}
	return
}

func (o *octo) increaseLevel() (flashers []*octo) {
	if !o.exhausted {
		o.energyLevel++

		if o.energyLevel == 10 {
			flashers = append(flashers, o)
		}
	}
	return
}

func (o *octo) addNeighbour(n *octo) {
	o.neighbours = append(o.neighbours, n)
	n.neighbours = append(n.neighbours, o)
}
func firstLine(read func() (string, bool)) (octos []*octo) {
	for ch, _ := read(); ch != "\n"; ch, _ = read() {
		el, _ := strconv.Atoi(ch)
		octo := octo{energyLevel: el}

		if x := len(octos); x != 0 {
			octos[x-1].addNeighbour(&octo)
		}

		octos = append(octos, &octo)
	}
	return
}
func parseOctos(input []byte) (octos []*octo) {
	read := utils.Reader(input, bufio.ScanBytes)
	previousLine := firstLine(read)
	octos = previousLine
	var currLine []*octo
	for ch, eoi := read(); !eoi; ch, eoi = read() {
		if ch == "\n" {
			octos = append(octos, currLine...)
			previousLine = currLine
			currLine = []*octo{}
			continue
		}
		el, _ := strconv.Atoi(ch)
		octo := octo{energyLevel: el}

		x := len(currLine)

		if x != 0 {
			currLine[x-1].addNeighbour(&octo)
			previousLine[x-1].addNeighbour(&octo)
		}

		previousLine[x].addNeighbour(&octo)
		currLine = append(currLine, &octo)

		if x < len(previousLine)-1 {
			previousLine[x+1].addNeighbour(&octo)
		}

	}
	octos = append(octos, currLine...)
	return
}

func iterate(octos []*octo) int {
	var flashers []*octo
	for _, o := range octos {
		flashers = append(flashers, o.increaseLevel()...)
	}

	var totalFlashers []*octo
	for _, o := range flashers {
		totalFlashers = append(totalFlashers, o.flash()...)
	}

	for _, o := range totalFlashers {
		o.exhausted = false
		o.energyLevel = 0
	}

	return len(totalFlashers)
}

func (DumboOctopus) SolveQ1(input []byte) int {
	octos := parseOctos(input)
	counter := 0
	for i := 0; i < 1000000; i++ {
		counter += iterate(octos)
	}
	return counter
}

func (DumboOctopus) SolveQ2(input []byte) int {
	counter := 0
	octos := parseOctos(input)
	for {
		counter++
		numberOfFlashes := iterate(octos)
		if numberOfFlashes == len(octos) {
			return counter
		}
	}
}
