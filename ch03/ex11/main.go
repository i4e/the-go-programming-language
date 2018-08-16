package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	n := len(s)

	if s[0:1] == "+" || s[0:1] == "-" {
		return s[0:1] + comma(s[1:])
	}

	dot := strings.LastIndex(s, ".")
	if dot >= 0 {
		return comma(s[:dot]) + s[dot:]
	}

	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaBuffer(s string) string {
	var buf bytes.Buffer
	firstComma := len(s) % 3

	if firstComma == 0 {
		firstComma = 3
	}
	buf.WriteString(s[:firstComma])

	for i := firstComma; i < len(s); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}
