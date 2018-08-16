// hamming distance

package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

var hashFunc = flag.String("hash", "sha256", "sha256, sha256, or sha256")

func main() {
	flag.Parse()

	in := bufio.NewReader(os.Stdin)
	for {
		bytes, err := in.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		switch *hashFunc {
		case "sha256":
			c := sha256.Sum256([]byte(bytes))
			fmt.Printf("%x\n", c)
		case "sha384":
			c := sha512.Sum384([]byte(bytes))
			fmt.Printf("%x\n", c)
		case "sha512":
			c := sha512.Sum512([]byte(bytes))
			fmt.Printf("%x\n", c)
		}
	}
}
