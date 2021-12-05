package day1

import (
	"strconv"
	"strings"
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
func solve(s string) int {
	lines := strings.Fields(s)
	var arr [2020]int

	for _, line := range lines {
		number, _ := strconv.Atoi(line)
		if arr[number] != -1 {
			arr[number] = number
		} else {
			return number * abs((number - 2020))
		}

		couple := abs(2020 - number)
		arr[couple] = -1
	}
	return -1
}
