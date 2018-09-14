package counter

import (
	"testing"
)

func TestLineCounter(t *testing.T) {
	tests := []struct {
		data     string
		expected int
	}{
		{
			"one",
			1,
		},
		{
			"one\ntwo\nthree\n",
			4,
		},
	}

	for _, test := range tests {
		var c LineCounter
		byteString := []byte(test.data)
		c.Write(byteString)
		if int(c) != test.expected {
			t.Errorf("expected %d, but actual %d", test.expected, c)
		}
	}
}

func TestWordCounter(t *testing.T) {
	tests := []struct {
		data     string
		expected int
	}{
		{
			"one",
			1,
		},
		{
			"one two three\nfour five six",
			6,
		},
	}

	for _, test := range tests {
		var c WordCounter
		byteString := []byte(test.data)
		c.Write(byteString)
		if int(c) != test.expected {
			t.Errorf("expected %d, but actual %d", test.expected, c)
		}
	}
}
