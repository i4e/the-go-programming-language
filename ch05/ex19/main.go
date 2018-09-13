package main

import (
	"fmt"
)

func main() {
	fmt.Println(f())
}

func f() (ret string) {
	defer func() {
		recover()
		ret = "ok"
	}()
	panic("test")
}
