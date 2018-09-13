package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"a", "b", "c", "d", "e"}
	s := joinMulArgs(" ", strs...)
	fmt.Println(s)
}

func joinMulArgs(sep string, strs ...string) string {
	combinedString := ""
	combinedString += strings.Join(strs, sep)
	return combinedString
}
