package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	m := make(map[string]int)
	tagFreq(doc, m)
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func tagFreq(n *html.Node, m map[string]int) {
	if n.Type == html.ElementNode {
		_, ok := m[n.Data]
		if ok {
			m[n.Data]++
		} else {
			m[n.Data] = 1
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		tagFreq(c, m)
	}
}
