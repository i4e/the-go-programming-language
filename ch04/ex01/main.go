// hamming distance

package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("a"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Println(c1[0])
	fmt.Println(HammingDistance(c1, c2))
	fmt.Println(HammingDistance2(c1, c2))
}

func HammingDistance(c1 [32]byte, c2 [32]byte) int {
	d := 0
	for i := 0; i < 32; i++ {
		// fmt.Println("d", d)
		x := c1[i]
		y := c2[i]
		d += int(x%2+y%2) % 2

		// fmt.Println(x, y)
		// fmt.Printf("%08b\n%08b\n", x, y)

		for i := 1; i < 8; i++ {
			x = x >> 1
			y = y >> 1
			// fmt.Println(x, y)
			// fmt.Printf("%08b\n%08b\n", x, y)
			d += int(x%2+y%2) % 2
		}
	}
	return d
}

func HammingDistance2(c1 [32]byte, c2 [32]byte) int {
	d := 0
	for i := 0; i < 32; i++ {
		// fmt.Println("d", d)
		diffBits := c1[i] ^ c2[i]
		d += PopCountClear(diffBits)
	}
	return d
}

func PopCountClear(x uint8) int {
	p := 0
	for x != 0 {
		x = x & (x - 1)
		p++
	}
	return p
}
