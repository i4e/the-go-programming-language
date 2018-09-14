package counter

import (
	"bytes"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	tests := []struct {
		data     string
		expected int64
	}{
		{
			"one",
			3,
		},
		{
			"one two three",
			13,
		},
	}

	for _, test := range tests {
		byteString := []byte(test.data)
		w, c := CountingWriter(bytes.NewBufferString(""))
		w.Write(byteString)
		if *c != test.expected {
			t.Errorf("expected %d, but actual %d", test.expected, c)
		}
	}
}
