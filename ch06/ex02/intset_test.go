package intset

import "fmt"

func Example_one() {
	var x IntSet
	x.AddAll(1, 2, 128, 192)

	//!+note
	fmt.Println(x.String()) // "{1 2 128 192}"
	fmt.Println(x)          // "{[3 0 65536]}"
	//!-note

	// Output:
	// {1 2 128 192}
	// {[6 0 1 1]}
}
