package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	// cyclic
	"linear algebra": {"calculus"},
	"hoge":           {"linear algebra"},
	"calculus":       {"hoge"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

type NodeStatus int

const (
	unvisited NodeStatus = iota
	visiting
	visited
)

func main() {
	sortedReqs, err := topoSort(prereqs)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, course := range sortedReqs {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	status := make(map[string]NodeStatus)
	var visitAll func(items []string) error

	visitAll = func(items []string) error {
		for _, item := range items {
			if status[item] == visiting {
				return fmt.Errorf("cyclic: %s", item)
			} else if status[item] == unvisited {
				status[item] = visiting
				err := visitAll(m[item])
				if err != nil {
					return err
				}
				status[item] = visited
				order = append(order, item)
			}
		}
		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	err := visitAll(keys)
	return order, err
}
