package day12

import (
	"adventofcode/utils"
	"bufio"
	"strings"
)

type PassagePathing struct{}

func (PassagePathing) Path() string {
	return "2021/12_passage_pathing"
}

type cave struct {
	name        string
	small       bool
	connections []*cave
}

func (c *cave) addConnectionTo(other *cave) {
	c.connections = append(c.connections, other)
	other.connections = append(other.connections, c)
}
func caveBuilder() (func(s string) *cave, map[string]*cave) {
	caves := map[string]*cave{}

	return func(s string) *cave {
		c, ok := caves[s]
		if !ok {
			c = &cave{name: s, small: strings.ToLower(s) == s}
			caves[c.name] = c
		}
		return c
	}, caves
}
func parseLine(line string) (cave1 string, cave2 string) {
	caves := strings.Split(line, "-")

	cave1 = caves[0]
	cave2 = caves[1]

	return
}
func prepareCaves(input []byte) map[string]*cave {
	newCave, caves := caveBuilder()
	read := utils.Reader(input, bufio.ScanLines)
	for line, eoi := read(); !eoi; line, eoi = read() {
		c1Name, c2Name := parseLine(line)
		cave1 := newCave(c1Name)
		cave2 := newCave(c2Name)

		cave1.addConnectionTo(cave2)
	}

	return caves
}

type path map[*cave]int

func (p path) copy() (fork path) {
	fork = make(path)
	for cave, visitCount := range p {
		fork[cave] = visitCount
	}
	return
}

var maxSmallVisit int

func (p path) canVisit(c *cave) bool {
	if !c.small {
		return true
	}

	if c.name == "start" && len(p) > 0 {
		return false
	}

	if _, ok := p[c]; ok {
		for c, v := range p {
			if c.small && v >= maxSmallVisit {
				return false
			}
		}
		return true
	}
	return true
}

func traverse(curr *cave, currPath path, allPaths *[]path) bool {
	if curr.name == "end" {
		return true
	}

	if !currPath.canVisit(curr) {
		return false
	}

	currPath[curr]++
	for _, next := range curr.connections {
		fork := currPath.copy()
		if traverse(next, fork, allPaths) {
			*allPaths = append(*allPaths, currPath)
		}
	}
	return false
}

func (PassagePathing) SolveQ1(input []byte) int {
	caves := prepareCaves(input)
	start := caves["start"]
	allPaths := &[]path{}

	maxSmallVisit = 1
	traverse(start, path{}, allPaths)

	return len(*allPaths)
}

func (PassagePathing) SolveQ2(input []byte) int {
	caves := prepareCaves(input)
	start := caves["start"]
	allPaths := &[]path{}

	maxSmallVisit = 2
	traverse(start, path{}, allPaths)

	return len(*allPaths)
}
