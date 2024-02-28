package main

import "fmt"

func main() {
	v := [2]string{"hello", "world"}
	fmt.Println(v)
	var v1 [2]int
	v1[0] = 1
	v1[1] = 9
	fmt.Println(v1)
}