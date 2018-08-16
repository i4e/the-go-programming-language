// hamming distance

package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("a"))
	c2 := sha256.Sum256([]byte("x"))
	fmt.Printf("%x\n%x\n", c1, c2)
	fmt.Println(HammingDistancePopCount(c1, c2))
}

func HammingDistance(c1 [32]byte, c2 [32]byte) int {
	d := 0
	for i := 0; i < 32; i++ {
		x := c1[i]
		y := c2[i]
		d += int(x%2+y%2) % 2

		for i := 1; i < 8; i++ {
			x = x >> 1
			y = y >> 1
			d += int(x%2+y%2) % 2
		}
	}
	return d
}

func HammingDistancePopCount(c1 [32]byte, c2 [32]byte) int {
	d := 0
	for i := 0; i < 32; i++ {
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
