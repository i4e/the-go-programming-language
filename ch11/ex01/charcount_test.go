package charcount

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCharCount(t *testing.T) {
	var tests = []struct {
		input   []byte
		counts  map[rune]int
		utflen  []int
		invalid int
	}{
		{
			input: nil,

			counts:  map[rune]int{},
			utflen:  []int{0, 0, 0, 0, 0},
			invalid: 0,
		},
		{
			input: []byte("a"),

			counts:  map[rune]int{'a': 1},
			utflen:  []int{0, 1, 0, 0, 0},
			invalid: 0,
		},
		{
			input: []byte("å"),

			counts:  map[rune]int{'å': 1},
			utflen:  []int{0, 0, 1, 0, 0},
			invalid: 0,
		},
		{
			input: []byte("あ"),

			counts:  map[rune]int{'あ': 1},
			utflen:  []int{0, 0, 0, 1, 0},
			invalid: 0,
		},
		{
			input: []byte{0xff},

			counts:  map[rune]int{},
			utflen:  []int{0, 0, 0, 0, 0},
			invalid: 1,
		},
		{
			input: []byte("aå"),

			counts:  map[rune]int{'a': 1, 'å': 1},
			utflen:  []int{0, 1, 1, 0, 0},
			invalid: 0,
		},
	}
	for _, test := range tests {
		counts, utflen, invalid, err := CharCount(bytes.NewReader(test.input))
		if err != nil {
			t.Error(err)
			continue
		}

		if got, expected := counts, test.counts; !reflect.DeepEqual(got, expected) {
			t.Errorf("want: %#v, got: %#v", expected, got)
		}
		if got, expected := utflen, test.utflen; !reflect.DeepEqual(got, expected) {
			t.Errorf("want: %#v, got: %#v", expected, got)
		}
		if got, expected := invalid, test.invalid; got != expected {
			t.Errorf("want: %#v, got: %#v", expected, got)
		}
	}
}
