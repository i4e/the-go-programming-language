package intset

import "fmt"

func Example_one() {
	var x, y IntSet
	x.AddAll(1, 2, 3, 4, 128, 192)
	y.AddAll(1, 2, 128, 129, 193)
	x.IntersectWith(&y)
	fmt.Println(x.String()) // "{1 2 128 192}"

	x.Clear()
	y.Clear()
	x.AddAll(1, 2, 3, 4, 128, 192)
	y.AddAll(1, 2, 128, 129, 193)
	x.DifferenceWith(&y)
	fmt.Println(x.String()) // "{1 2 128 192}"

	x.Clear()
	y.Clear()
	x.AddAll(1, 2, 3, 4, 128, 192)
	y.AddAll(1, 2, 128, 129, 193)
	x.SymmetricDifference(&y)
	fmt.Println(x.String()) // "{1 2 128 192}"

	// Output:
	// {1 2 128}
	// {3 4 192}
	// {3 4 129 192 193}
}
