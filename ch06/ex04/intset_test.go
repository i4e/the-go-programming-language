package intset

import "fmt"

func Example_one() {
	var x IntSet
	x.AddAll(1, 2, 3, 4, 128, 192)
	fmt.Println(x.Elems())

	// Output:
	// [1 2 3 4 128 192]
}
