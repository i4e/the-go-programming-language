package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

func main() {
	fmt.Println(isAnagram(os.Args[1], os.Args[2]))
}

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	aList := strings.Split(a, "")
	bList := strings.Split(b, "")
	sort.Strings(aList)
	sort.Strings(bList)
	return reflect.DeepEqual(aList, bList)
}
