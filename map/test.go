package main

import (
	"fmt"
	"strings"
)

func main() {
	var t = strings.Fields("  foo for bar  foo  baz   ")
	m := make(map[string]int)
	for _, v := range t {
		_, ok := m[v]
		if ok {
			m[v] += 1
		} else {
			m[v] = 1

		}
	}
	fmt.Println(m)
}