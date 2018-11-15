package charcount

import (
	"bufio"
	"io"
	"unicode"
	"unicode/utf8"
)

func CharCount(r io.Reader) (map[rune]int, []int, int, error) {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(r)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			return counts, utflen[:], invalid, nil
		}
		if err != nil {
			return nil, nil, 0, err
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
}
