package main

import "fmt"

type T interface {
	M()
}

type K struct {
	i int
}

func (k *K) M() {
	if k == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(k.i)
}

func main() {
	var k T
	var t *K
	k = t
	k.M()
}