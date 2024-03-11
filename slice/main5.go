package main

import "fmt"

func main() {
	a := make([]int, 7)
	fmt.Println(a, len(a), cap(a))
	b := make([]string, 7)
	fmt.Println(b, len(b), cap(b))
	c := make([]int, 2, 7)
	fmt.Println(c, len(c), cap(c), c[0])
}