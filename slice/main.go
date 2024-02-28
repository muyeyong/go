package main

import "fmt"

func main() {
	v := [3]int{1, 2, 3}
	v1 := v[1:2]
	fmt.Println(v, v1)
}