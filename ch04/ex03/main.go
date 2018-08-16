package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(a)
	reverse(a[:])
	fmt.Println(a)
	b := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(&b)
	reverseWithPointer(&b)
	fmt.Println(&b)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseWithPointer(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
