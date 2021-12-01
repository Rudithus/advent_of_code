package day1

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve(s string) int {
	lines := strings.Fields(s)

	counter := 0
	previous, _ := strconv.ParseInt(lines[0], 10, 32)
	for i := 1; i < len(lines); i++ {
		depth, _ := strconv.ParseInt(lines[i], 10, 32)

		if depth > previous {
			counter++
		}

		previous = depth
	}

	fmt.Println(counter)
	return counter
}
