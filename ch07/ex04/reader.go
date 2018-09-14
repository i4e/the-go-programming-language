package reader

import (
	"io"
)

type stringReader string

func (s stringReader) Read(p []byte) (n int, err error) {
	n = copy(p, s)
	s = s[n:]
	if len(s) == 0 {
		err = io.EOF
	}
	return
}

func NewReader(s string) stringReader {
	return stringReader(s)
}
