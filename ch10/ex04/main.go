package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Package struct {
	ImportPath string   // import path of package in dir
	Name       string   // package name
	Deps       []string // all (recursively) imported dependencies
}

func main() {
	var err error
	key := os.Args[1]
	cmd := exec.Command("go", "list", key)
	if _, err = cmd.Output(); err != nil {
		log.Fatalf("package %s invalid: %v", key, err)
	}

	cmd = exec.Command("go", "list", "-json", "...")
	if cmd == nil {
		log.Fatalf("can't run go list")
	}

	var output []byte
	if output, err = cmd.Output(); err != nil {
		log.Fatal(err)
	}

	var stack []byte
	var buf bytes.Buffer
	for _, b := range output {
		switch b {
		case '{':
			stack = append(stack, b)
		case '}':
			stack = stack[0 : len(stack)-1]
		}

		buf.WriteByte(b)
		if b == '}' && len(stack) == 0 {
			var info Package
			if err = json.Unmarshal(buf.Bytes(), &info); err != nil {
				log.Fatal(err)
			}
			if contain(info.Deps, key) {
				fmt.Println(info.ImportPath)
			}
			buf.Truncate(0)
		}
	}
}

func contain(array []string, key string) bool {
	low, high := 0, len(array)-1
	for low <= high {
		var mid = (low + high) / 2
		if key == array[mid] {
			return true
		}
		if key < array[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
