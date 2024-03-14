package main

import (
	"fmt"
	"io"
	"strings"
)

type MyReader struct {
	S string
}

func (s MyReader) Read() {
	r := strings.NewReader(s.S)
	b := make([]byte, 1)
	for {
		n, err := r.Read(b)
		fmt.Printf("%q", b[:n])
		if err == io.EOF {
			r.Reset(s.S)
		}
	}
}

func main() {
	r := MyReader{S: "A"}
	r.Read()
}