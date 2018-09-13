package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var s string
	if sc.Scan() {
		s = sc.Text()
	}
	fmt.Println(expand(s, replace))
}

func expand(s string, f func(string) string) string {
	w := func(str string) string {
		return f(str[1:])
	}

	var retvals []string

	for _, v := range strings.Split(s, " ") {
		if v[0] == '$' {
			retvals = append(retvals, w(v))
		} else {
			retvals = append(retvals, v)
		}
	}
	return strings.Join(retvals, " ")
}

func replace(s string) string {
	return "doller_" + s
}
