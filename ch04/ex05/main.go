package main

import "fmt"

func main() {
	a := []string{"a", "b", "b", "c", "c", "c", "c", "d", "e"}
	fmt.Println(a)
	a = dleteAdjacent(a)
	fmt.Println(a)
}

func dleteAdjacent(strings []string) []string {
	ptr := 0
	for i := 0; i < len(strings)-1; i++ {
		if strings[ptr] == strings[i] {
			continue
		}
		ptr++
		strings[ptr] = strings[i]
	}
	return strings[:ptr+1]
}
