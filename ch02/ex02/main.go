package main

import (
	"bufio"
	"fmt"
	"os"

	"strconv"
	"strings"

	"./tempconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	if len(os.Args) == 1 {
		var s string
		if sc.Scan() {
			s = sc.Text()
		}
		for _, s := range strings.Split(s, " ") {
			t, err := strconv.ParseFloat(s, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			hoge(t)
		}
	} else {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(1)
			}
			hoge(t)
		}
	}

}

func hoge(t float64) {
	c := tempconv.Celsius(t)
	f := tempconv.Fahrenheit(t)
	k := tempconv.Kelvin(t)
	fmt.Println(t)
	fmt.Println(tempconv.CToF(c))
	fmt.Println(tempconv.CToK(c))
	fmt.Println(tempconv.KToF(k))
	fmt.Println(tempconv.KToC(k))
	fmt.Println(tempconv.FToC(f))
	fmt.Println(tempconv.FToK(f))
}
