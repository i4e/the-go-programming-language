package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FileCount struct {
	filename []string
	count    int
}

func main() {
	counts := make(map[string]*FileCount)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, counts := range counts {
		filenames := strings.Join(counts.filename, " ")
		if counts.count > 1 {
			fmt.Printf("%d %s\n  %v\n", counts.count, line, filenames)
		}
	}
}

func in(needle string, strings []string) bool {
	for _, s := range strings {
		if needle == s {
			return true
		}
	}
	return false
}

func countLines(f *os.File, counts map[string]*FileCount) {
	input := bufio.NewScanner(f)
	filename := f.Name()

	for input.Scan() {
		line := input.Text()
		if c, ok := counts[line]; ok {
			c.count++
			if !in(filename, c.filename) {
				c.filename = append(c.filename, filename)
			}
		} else {
			counts[line] = &FileCount{[]string{filename}, 1}
		}
	}
}
