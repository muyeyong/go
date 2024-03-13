package main

import (
	"fmt"
)

type T interface {
	M()
}

type K struct {
	i string
}

func (k K) M() {
	fmt.Println(k.i)
}

func main() {
	var t T = K{"hello"}
	t.M()
	
}