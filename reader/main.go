package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, World!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("%q", b[:n])
		fmt.Print(string(b[:n]))
		fmt.Print(string(b))
		if err == io.EOF {
			break
		}
	}
}