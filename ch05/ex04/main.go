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
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, "a "+a.Val)
				}
			}
		case "img":
			for _, a := range n.Attr {
				if a.Key == "src" {
					links = append(links, "img "+a.Val)
				}
			}

		case "script":
			for _, a := range n.Attr {
				if a.Key == "src" {
					links = append(links, "script "+a.Val)
				}
			}

		case "link":
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, "link "+a.Val)
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
