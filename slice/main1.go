package main

import "fmt"

func main() {
	var v [3]int
	v[0] = 1
	fmt.Println(v)
	v1 := [3]string{"1", "5", "Monday"}
	fmt.Println(v1)
	v11 := v1[1:2]
	fmt.Println(v11)
	v1[1] = "Tuesday"
	fmt.Println(v11)
}