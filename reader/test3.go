package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read (b []byte) (int, error) {
	n, err := r13.r.Read(b)
	for i := 0; i < n; i++ {
		c := b[i]
		if 'A' <= c && c <= 'M' || 'a' <= c && c <= 'm' {
			b[i] += 13
		} else if 'N' <= c && c <= 'Z' || 'n' <= c && c <= 'z' {
			b[i] -= 13
		}
	}
	return n, err

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}