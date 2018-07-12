package main

import (
	"fmt"

	"./tempconv"
)

func main() {
	const (
		c tempconv.Celsius    = 100.0
		f tempconv.Fahrenheit = 100.0
		k tempconv.Kelvin     = 100.0
	)

	fmt.Println(tempconv.CToF(c))
	fmt.Println(tempconv.CToK(c))
	fmt.Println(tempconv.KToF(k))
	fmt.Println(tempconv.KToC(k))
	fmt.Println(tempconv.FToC(f))
	fmt.Println(tempconv.FToK(f))
}
