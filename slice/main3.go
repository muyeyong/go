package main

import "fmt"

func main() {
	v := []int{1, 2, 3, 4, 5, 6}
	v1 := v[:]
	v2 := v[0:]
	v3 := v[:5]
	v4 := v[0:5]
	fmt.Println(v1, v2, v3, v4)
}