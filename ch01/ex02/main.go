package main

import (
	"fmt"
	"os"
)

func main() {
	echo2()
}

func echo1() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}

}

func echo2() {
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
}
