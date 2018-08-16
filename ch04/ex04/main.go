package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(a)
	rotate(a, 3)
	fmt.Println(a)
}

func rotate(ints []int, n int) {
	copy(ints, append(ints[n:], ints[:n]...))
}

func rotate2(ints []int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < len(ints)-1; j++ {
			ints[j], ints[j+1] = ints[j+1], ints[j]
		}
	}
}
