package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil !")
	}
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
}