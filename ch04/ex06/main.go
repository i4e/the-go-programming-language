package main

import (
	"fmt"
	"unicode"
)

func main() {
	a := "aaaaa aaa    aa  aaaaaaa"
	b := []byte(a)
	c := removeDuplicatedSpaces(b)
	fmt.Println(string(c))
}

// FIXME: 「もとのスライス」
func removeDuplicatedSpaces(bs []byte) []byte {
	result := bs[:0]
	for i, b := range bs {
		if unicode.IsSpace(rune(b)) {
			if i > 0 && unicode.IsSpace(rune(bs[i-1])) {
				continue
			} else {
				result = append(result, ' ')
			}
		} else {
			result = append(result, b)
		}
	}
	return result
}
