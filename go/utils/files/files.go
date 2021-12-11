package files

import (
	"adventofcode/utils"
	"bufio"
	"os"
	"strings"
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

func commaSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := strings.Index(string(data), ","); i >= 0 {
		return i + 1, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return
}

func CharReader(path string) utils.ReadLine {
	file, err := os.Open(path)
	utils.Check(err)

	scanner := bufio.NewScanner(bufio.NewReader(file))
	scanner.Split(commaSplitFunc)
	return func() (line string, eof bool) {
		eof = !scanner.Scan()

		return scanner.Text(), eof
	}
}
