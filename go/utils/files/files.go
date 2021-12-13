package files

import (
	"adventofcode/utils"
	"bufio"
	"os"
)

type FileReader func() (line string, eoi bool)

func LineReader(path string) FileReader {
	file, err := os.Open(path)
	utils.Check(err)
	scanner := bufio.NewScanner(bufio.NewReader(file))
	return func() (line string, eof bool) {
		eof = !scanner.Scan()

		return scanner.Text(), eof
	}
}

func CharReader(path string) FileReader {
	file, err := os.Open(path)
	utils.Check(err)

	scanner := bufio.NewScanner(bufio.NewReader(file))
	scanner.Split(utils.CommaSplitFunc)
	return func() (line string, eof bool) {
		eof = !scanner.Scan()

		return scanner.Text(), eof
	}
}
