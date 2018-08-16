package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	a := "abcdefgあいうえお"
	b := []byte(a)
	reverseBytes(b)
	fmt.Println(string(b))
}

func reverseBytes(b []byte) {
	var size int
	for i := 0; i < len(b); i += size {
		_, size = utf8.DecodeRune(b[i:])
		reverse(b[i : i+size])
	}
	reverse(b)
}

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
