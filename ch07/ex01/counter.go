package counter

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type LineCounter int

func (c *LineCounter) Write(p []byte) (n int, err error) {
	*c++
	for _, b := range p {
		if b == '\n' {
			*c++
		}
	}
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (n int, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c++
	}
	if err = scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	return len(p), err
}
