package files

import (
	"adventofcode/utils"
	"bufio"
	"os"
)

func LineReader(path string) utils.ReadLine {
	file, err := os.Open(path)
	utils.Check(err)

	scanner := bufio.NewScanner(bufio.NewReader(file))
	return func() (line string, eof bool) {
		eof = !scanner.Scan()

		return scanner.Text(), eof
	}
}
