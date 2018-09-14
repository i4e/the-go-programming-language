package counter

import (
	"io"
)

type Counter struct {
	writer io.Writer
	count  int64
}

func (c *Counter) Write(p []byte) (n int, err error) {
	c.count += int64(len(p))
	n, err = c.writer.Write(p)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &Counter{w, 0}
	return c, &c.count
}
