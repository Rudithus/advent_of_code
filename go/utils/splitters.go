package utils

import (
	"bufio"
	"bytes"
	"strings"
)

func CommaSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {

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

func Reader(byteArr []byte, splitter bufio.SplitFunc) func() (line string, eoi bool) {
	scanner := bufio.NewScanner(bytes.NewReader(byteArr))
	scanner.Split(splitter)

	return func() (line string, eof bool) {
		eof = !scanner.Scan()

		return scanner.Text(), eof
	}
}
