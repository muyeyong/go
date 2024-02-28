package main

import "fmt"

func main() {
	i := 34
	p := &i
	fmt.Println(i)
	fmt.Println(*p)
}